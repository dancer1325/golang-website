//go:build OMIT

package main

import "fmt"

func main() {
	// *string			type of the pointer / value's type == string
	var d *string
	// zero value of a pointer == nil   -- Reason: NOT declared --
	fmt.Println("d", d)

	// types inferred
	i, j := 42, 2701

	// &operand -- generates a -> pointer of the operand
	p := &i // -- generates a -> pointer of i /  holds the memory address of i
	fmt.Println("p", p)
	fmt.Println("*p", *p) // read i -- through the -- pointer
	// pointers are NOT arithmetic
	//p++

	// *pointerValue = value
	*p = 21 // p is a pointer -> *p is the value i
	fmt.Println("i", i)
	fmt.Println("*p", *p) // same value as before

	p = &j // -- generates a -> pointer of j / holds the memory address of j -> another value now for p
	fmt.Println("p", p)
	fmt.Println("i", i) // i has NOT changed
	fmt.Println("j", j)
	*p = *p / 37        // p is a pointer -> *p is the value j which is divided through the pointer
	fmt.Println("p", p) // Although j changes, pointer does NOT change
	fmt.Println("j", j)

	// Just declare == zero values
	var a string
	b := &a // There is already a memory address booked for that variable, although it's just declared!!
	fmt.Println("a", a)
	fmt.Println("b", b)

}
