package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// a := ex1()
	// ex2(a)
	ex3()
}

// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.
func ex1() []int {
	fmt.Println("Exercise 1")
	a := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		a = append(a, rand.Intn(100))
	}

	fmt.Println("a:", a)
	fmt.Println()
	return a
}

// Loop over the slice you created in exercise 1. For each value in the slice, apply the following rules:
// a. If the value is divisible by 2, print “Two!”
// b. If the value is divisible by 3, print “Three!”
// c. If the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
// d. Otherwise, print “Never mind”.
func ex2(x []int) {
	fmt.Println("Exercise 2")

	for _, v := range x {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
	fmt.Println()

}

// Declare an int variable called total.
// Write a for loop that uses a variable named i to iterate from 0 (inclusive) to 10 (exclusive).
func ex3() {
	fmt.Println("Exercise 3")
	var total int

	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}

	fmt.Println(total)
	// This prints out 0; we have shadowed the function block var int in the for block, so we are not actually incrementing the correct var.
	// Each iteration, the new var total is simply assigned the value of the iterator 'i'

	// What we should do:
	for i := 0; i < 10; i++ {
		total = total + i
		fmt.Println(total)
	}

	fmt.Println(total)

}
