package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	// struct generic type passing bool
	var v1 = List[bool]{}
	fmt.Println(v1.val)
}
