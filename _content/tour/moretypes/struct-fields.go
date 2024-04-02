//go:build OMIT
// +build OMIT

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println("v", v)

	// `structVariable.structField`		--	get access to the — fields
	v.X = 4 // set a new value for the field
	fmt.Println("v.X", v.X)
	fmt.Println("v", v)
}
