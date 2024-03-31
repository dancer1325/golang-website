//go:build OMIT
// +build OMIT

package main

import "fmt"

// (x, y int)		-- named return values --
func split(sum int) (x, y int) {
	// named return values 		== variables declared as top of the function
	x = sum * 4 / 9
	y = sum - x
	return // "naked" return  == return the named return values
}

func main() {
	fmt.Println(split(17))
}
