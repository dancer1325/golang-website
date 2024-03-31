//go:build OMIT
// +build OMIT

package main

import "fmt"

// Initialization   -- 1/variable --
var i, j int = 1, 2

func main() {
	// If you initialize -> type declaration can be omitted
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
