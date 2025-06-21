//go:build OMIT

package main

import "fmt"

func main() {
	// make([]T, lengthValue, capacityValue)		--	create a slice
	pow := make([]int, 10)

	// i 	-- slice item’s position
	// second possible return value is omitted
	for i := range pow {
		fmt.Println(fmt.Sprintf("i - %v, uint()i %v", i, uint(i)))
		pow[i] = 1 << uint(i) // == 2**i
		fmt.Println(fmt.Sprintf("pow[i] %v", pow[i]))
	}

	// _	-- slice item's position is skipped
	// v	-- copy of the slice item’s value
	for _, value := range pow {
		// %d		integer variable -- is formatted to -> decimal number
		fmt.Printf("%d\n", value)
	}
}
