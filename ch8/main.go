package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
}

// Write a generic function that doubles the value of any integer or float that’s passed in to it.
// Define any needed generic interfaces.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Doubler[T Number](t T) T {
	return t * 2
}

func ex1() {
	fmt.Println("Exercise 1")
	fmt.Println(Doubler(16.2))
	fmt.Println(Doubler(35))
	fmt.Println("")
}

// Define a generic interface called Printable that matches a type that implements fmt.Stringer and has an underlying type of int or float64.
// Define types that meet this interface.
// Write a function that takes in a Printable and prints its value to the screen using fmt.Println.
type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintInt int
type PrintFloat float64

func (pi PrintInt) String() string {
	return fmt.Sprintf("%d", pi)

}

func (pf PrintFloat) String() string {
	return fmt.Sprintf("%f", pf)
}

func PrintNumber[T Printable](t T) {
	fmt.Println(t)
}

func ex2() {
	fmt.Println("Exercise 2")
	var i PrintInt = 5
	PrintNumber(i)

	var f PrintFloat = 3.543
	PrintNumber(f)
	fmt.Println()
}

// Write a generic singly linked list data type.
// Each element can hold a comparable value and has a pointer to the next element in the list.

type SinglyLinkedList[T comparable] struct {
	Head *Element[T]
	Tail *Element[T]
}

type Element[T comparable] struct {
	val  T
	next *Element[T]
}

// adds a new element to the end of the linked list
func (l *SinglyLinkedList[T]) Add(t T) {
	e := &Element[T]{val: t}

	// if list empty, set new el as head and tail
	if l.Head == nil {
		l.Head = e
		l.Tail = e
		return
	}

	// update current tail
	l.Tail.next = e

	// set new tail
	l.Tail = e
}

// adds an element at the specified position in the linked list
func (l *SinglyLinkedList[T]) Insert(t T, p int) {
	e := &Element[T]{val: t}

	// if list empty, set new el as head and tail (ignore position)
	if l.Head == nil {
		l.Head = e
		l.Tail = e
		return
	}

	// insert as head
	if p <= 0 {
		e.next = l.Head
		l.Head = e
		return
	}

	curr := l.Head
	// iterate over list
	for i := 1; i < p; i++ {
		// if at end, add new el as tail
		if curr.next == nil {
			curr.next = e
			l.Tail = e
			return
		}
		curr = curr.next
	}
	e.next = curr.next
	curr.next = e

	if l.Tail == curr {
		l.Tail = e
	}
}

// returns the position of the supplied value, -1 if it's not present
func (l *SinglyLinkedList[T]) Index(t T) int {
	i := 0

	for curr := l.Head; curr != nil; curr = curr.next {
		if curr.val == t {
			return i
		}
		i++
	}

	return i
}

func ex3() {
	fmt.Println("Exercise 3")

	list := &SinglyLinkedList[int]{}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	fmt.Println(list.Index(5))
	fmt.Println(list.Index(3))

	list.Insert(99, 0)
	fmt.Println(list.Index(99))
	fmt.Println(list.Index(5))

	fmt.Println()
}
