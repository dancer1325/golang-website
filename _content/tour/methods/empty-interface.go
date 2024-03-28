//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	// interface{}  can hold values of any type
	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

// interface{}  because `Printf` can handle values of any type
func describe(i interface{}) {
	// %v		-- prints the -- operand's value
	// %T		-- prints the -- operand's type
	fmt.Printf("(%v, %T)\n", i, i)
}
