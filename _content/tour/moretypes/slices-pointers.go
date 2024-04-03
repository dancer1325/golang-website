//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	// array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// create slices based on the previous array
	a := names[0:2] // 0 is the first element		& highestIndex is NOT included
	b := names[1:3]
	fmt.Println(a, b)

	// modify ONE of the slice's elements
	b[0] = "XXX"
	fmt.Println(a, b)  // both slices are affected
	fmt.Println(names) // underlying array is affected
}
