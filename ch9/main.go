package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Sentinel Error
var ErrInvalidId = errors.New("invalid id")

type EmptyFieldErr struct {
	FieldName string
	Err       error
}

func (ef EmptyFieldErr) Error() string {
	return ef.FieldName
}

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		message := fmt.Sprintf("record %d: %+v", count, emp)

		if err != nil {
			message = message + " error(s): "
			var allErrs []error
			if errs, ok := err.(interface{ Unwrap() []error }); ok {
				allErrs = errs.Unwrap()
			}
			for _, e := range allErrs {
				message = message + ProcessErr(e, emp) + " "
			}
		}
		fmt.Println(message)

	}
}

func ProcessErr(err error, emp Employee) string {
	var fieldErr EmptyFieldErr
	if errors.Is(err, ErrInvalidId) {
		return fmt.Sprintf("invalid ID: %s", emp.ID)
	} else if errors.As(err, &fieldErr) {
		return fmt.Sprintf("empty field: %s", fieldErr.FieldName)
	} else {
		return fmt.Sprintf("%v", err)
	}
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var (
	validID = regexp.MustCompile(`\w{4}-\d{3}`)
)

func ValidateEmployee(e Employee) error {
	var errs []error
	if len(e.ID) == 0 {
		errs = append(errs, EmptyFieldErr{FieldName: "ID"})
	}
	if !validID.MatchString(e.ID) {
		errs = append(errs, ErrInvalidId)
	}
	if len(e.FirstName) == 0 {
		errs = append(errs, EmptyFieldErr{FieldName: "FirstName"})
	}
	if len(e.LastName) == 0 {
		errs = append(errs, EmptyFieldErr{FieldName: "LastName"})
	}
	if len(e.Title) == 0 {
		errs = append(errs, EmptyFieldErr{FieldName: "Title"})
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
