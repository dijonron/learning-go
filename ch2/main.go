package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
}

// Write a program that declares an integer variable called i with the value 20.
// Assign i to a floating-point variable named f. Print out i and f.
func ex1() {
	fmt.Println("Exercise 1")
	i := 20
	f := float64(i) // we don't need to specify the type of f, but we do need to explicitly convert it to float64

	fmt.Println("i:", i)
	fmt.Println("f:", f)
	fmt.Println()

}

// Write a program that declares a constant called value that can be assigned to both an integer and a floating-point variable.
// Assign it to an integer called i and a floating-point variable called f. Print out i and f.
func ex2() {
	fmt.Println("Exercise 2")
	const value = 10

	i := value            // since value is an int literal, we do not need to specify the type for i
	var f float64 = value // since value is an int literal, we need to specify the type for f

	fmt.Println("i:", i)
	fmt.Println("f:", f)
	fmt.Println()
}

// Write a program with three variables, one named b of type byte, one named smallI of type int32, and one named bigI of type uint64.
// Assign each variable the maximum legal value for its type; then add 1 to each variable. Print out their values.
func ex3() {
	fmt.Println("Exercise 3")
	var b byte = 255                       // max value for byte
	var smallI int32 = 2147483647          // max value for int32
	var bigI uint64 = 18446744073709551615 // max value for uint64

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println("b:", b)
	fmt.Println("smallI:", smallI)
	fmt.Println("bigI:", bigI)
	fmt.Println()

	// causes overflow to the minimum value for each type, not an error
}
