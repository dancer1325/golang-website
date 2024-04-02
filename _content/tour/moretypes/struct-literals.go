//go:build OMIT
// +build OMIT

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	// Without specifying fields -> by order
	v1 = Vertex{1, 2}

	// Specifying fields
	v2 = Vertex{Y: 1} // X:0 is implicit

	v3 = Vertex{} // X:0 and Y:0

	// Pointer to the structVariable
	p = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
