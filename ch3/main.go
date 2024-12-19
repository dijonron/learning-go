package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
}

// Write a program that defines a variable named greetings of type slice of strings with the following values: "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт".
// Create a subslice containing the first two values; a second subslice with the second, third, and fourth values; and a third subslice with the fourth and fifth values.
// Print out all four slices.
func ex1() {
	fmt.Println("Exercise 1")
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	ss1 := greetings[:2]
	ss2 := greetings[1:4]
	ss3 := greetings[3:]

	fmt.Println("greetings:", greetings)
	fmt.Println("ss1:", ss1)
	fmt.Println("ss2:", ss2)
	fmt.Println("ss3:", ss3)
	fmt.Println()
}

// Write a program that defines a string variable called message with the value "Hi 👩 and 👨" and prints the fourth rune in it as a character, not a number.
func ex2() {
	fmt.Println("Exercise 2")
	var message string = "Hi 👩 and 👨"

	var character rune = []rune(message)[3] // convert string to array of runes, slice to get 4th rune

	fmt.Println("message:", message)
	fmt.Println("character:", string(character)) // convert rune to string and print
	fmt.Println()
}

// Write a program that defines a struct called Employee with three fields: firstName, lastName, and id.
// The first two fields are of type string, and the last field (id) is of type int.
// Create three instances of this struct using whatever values you’d like.
// Initialize the first one using the struct literal style without names, the second using the struct literal style with names, and the third with a var declaration. Use dot notation to populate the fields in the third struct.
// Print out all three structs.
func ex3() {
	fmt.Println("Exercise 3")

	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	employee1 := Employee{"Dalton", "Ronan", 1}

	employee2 := Employee{
		firstName: "Homer",
		lastName:  "Simpson",
		id:        2,
	}

	var employee3 Employee
	employee3.firstName = "Ronald"
	employee3.lastName = "McDonald"
	employee3.id = 3

	fmt.Println("employee1:", employee1)
	fmt.Println("employee2:", employee2)
	fmt.Println("employee3:", employee3)
}
