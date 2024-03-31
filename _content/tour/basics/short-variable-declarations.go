//go:build OMIT
// +build OMIT

package main

import "fmt"

// Short declaration NOT valid outside a function
//a := true

func main() {
	var i, j int = 1, 2

	// short declaration
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
