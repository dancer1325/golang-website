//go:build OMIT
// +build OMIT

package main

import "fmt"

// slice literal		-- passing the items values directly
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// i 	-- slice item’s position
	// v	-- copy of the slice item’s value
	for i, v := range pow {
		// %d		integer variable -- is formatted to -> decimal number
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
