//go:build OMIT

package main

import "fmt"

func main() {
	// make([]T, lengthValue, capacityValue)
	// capacityValue NOT specified -> capacityValue = lengthValue
	a := make([]int, 5)
	printSlice("a", a)

	// make([]T, lengthValue, capacityValue)
	b := make([]int, 0, 5)
	printSlice("b", b)

	// Reassign the slice adjusting the highestIndex ->
	// 1. slice’s length changes
	// 2. array NOT change → capacity NOT change
	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	// %s		-- string value, representing the variable name
	// %d		-- signed integers
	// %v		-- value in default format
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
