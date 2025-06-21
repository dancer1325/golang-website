//go:build OMIT

package main

import "fmt"

func main() {
	// slice literal
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Reassign the slice adjusting the highestIndex ->
	// 1. slice’s length changes
	// 2. array NOT change → capacity NOT change
	s = s[:0]
	printSlice(s)

	// Reassign the slice adjusting the highestIndex ->
	// 1. slice’s length changes
	// 2. array NOT change → capacity NOT change
	s = s[:4]
	printSlice(s)

	// Reassign the slice adjusting the lowestIndex ->
	// 1. slice’s length changes &
	// 2. array changes / drops elements -> capacity changes
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	// %d		-- signed integer --
	// %v		-- value in default format --
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
