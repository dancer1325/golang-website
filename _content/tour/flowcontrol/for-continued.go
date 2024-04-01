//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	sum := 1

	// for initStatement; condition; postStatement {...}
	// initStatement			-- optional -- NOT in the next loop
	// condition				-- mandatory --
	// postStatement			-- optional -- NOT in the next loop
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
