package main

import "fmt"

func pointers() {
	x := 10
	pointerToX := &x         // holds the address of x
	fmt.Println(pointerToX)  // prints mem add. eg  0x0001
	fmt.Println(*pointerToX) // dereferences pointerToX, i.e. prints value of x, i.e. prints 10
	z := 5 + *pointerToX     // creates new var z which is 5 plus dereferenced pointerToX, i.e. 5 + 10
	fmt.Println(z)           // prints 15
	fmt.Println()
}
