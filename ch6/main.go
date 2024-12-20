package main

import "fmt"

func main() {
	// pointers()

	// ex1()
	// ex2()
	ex3()
}

// Create a struct named Person with three fields: FirstName and LastName of type string and Age of type int.
// Write a function called MakePerson that takes in firstName, lastName, and age and returns a Person.
// Write a second function MakePersonPointer that takes in firstName, lastName, and age and returns a *Person.
// Call both from main.
// Compile your program with go build -gcflags="-m". This both compiles your code and prints out which values escape to the heap. Are you surprised about what escapes?
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

}

func ex1() {
	fmt.Println("Exercise 1")

	p := MakePerson("Dalton", "Ronan", 30)
	fmt.Println("person:", p)

	pp := MakePersonPointer("Ronald", "McDonald", 84)
	fmt.Println("person pointer:", pp)

	//	./main.go:46:25: &Person{...} escapes to heap
	// the pointer is returned to the stack, but he value it points to escapes to the heap
	fmt.Println()
}

// The UpdateSlice function takes in a []string and a string.
// It sets the last position in the passed-in slice to the passed-in string.
// At the end of UpdateSlice, print the slice after making the change.
func UpdateSlice(s []string, v string) {
	s[len(s)-1] = v
	fmt.Println("UpdateSlice:", s)
}

// The GrowSlice function also takes in a []string and a string.
// It appends the string onto the slice.
// At the end of GrowSlice, print the slice after making the change.
func GrowSlice(s []string, v string) {
	s = append(s, v)
	fmt.Println("GrowSlice:", s)
}

// Call these functions from main.
// Print out the slice before each function is called and after each function is called.
// Do you understand why some changes are visible in main and why some changes are not?
func ex2() {
	fmt.Println("Exercise 2")
	s := []string{"a", "b", "c", "d"}
	fmt.Println("original slice:", s)

	UpdateSlice(s, "...z")
	fmt.Println("updated slice:", s)

	GrowSlice(s, "e")
	fmt.Println("grown slice:", s)
	fmt.Println("")

	// Output:
	// original slice: [a b c d]
	// UpdateSlice: [a b c ...z]
	// updated slice: [a b c ...z]
	// GrowSlice: [a b c ...z e]
	// grown slice: [a b c ...z]

	// The update slice is reflected in main because wwe are updating the contents of the slice; since the copy and original pointer of s points to the update, we see it
	// The grown slice is not reflected in main because the only len of the copy is updated; the original slice's len is not changed
}

// Write a program that builds a []Person with 10,000,000 entries (they could all be the same names and ages).
// See how long it takes to run.
// Change the value of GOGC and see how that affects the time it takes for the program to complete.
// Set the environment variable GODEBUG=gctrace=1 to see when garbage collections happen and see how changing GOGC changes the number of garbage collections.
// What happens if you create the slice with a capacity of 10,000,000?
func ex3() {
	fmt.Println("Exercise 3")
	const numPeople = 10_000_000

	// do not pre-allocate slice
	// var people []Person

	// pre-allocate slice
	people := make([]Person, 0, numPeople)

	for i := 0; i < numPeople; i++ {
		people = append(people, MakePerson("John", "Doe", 35))
	}

	// Results (other ex's commented out before executing):
	// no allocation
	// GOGC=100 ./ch_6  1.45s user 0.19s system 145% cpu 1.126 total
	// GOGC=50 ./ch_6  1.64s user 0.25s system 195% cpu 0.965 total
	// GOGC=200 ./ch_6  0.87s user 0.14s system 191% cpu 0.527 total
	// GOGC=1000 ./ch_6  0.44s user 0.22s system 143% cpu 0.463 total
	// GOGC=off ./ch_6  0.13s user 0.24s system 82% cpu 0.442 total

	// pre-allocated slice
	// GOGC=100 ./ch_6  0.11s user 0.08s system 64% cpu 0.300 total
	// GOGC=50 ./ch_6  0.14s user 0.14s system 351% cpu 0.077 total
	// GOGC=200 ./ch_6  0.14s user 0.14s system 343% cpu 0.081 total
	// GOGC=1000 ./ch_6  0.15s user 0.12s system 316% cpu 0.086 total
	// GOGC=off ./ch_6  0.03s user 0.03s system 92% cpu 0.067 total
}
