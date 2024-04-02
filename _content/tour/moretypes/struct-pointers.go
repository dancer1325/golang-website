//go:build OMIT
// +build OMIT

package main

import "fmt"

// type structName struct { fields }
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	// pointer to the struct variable
	p := &v

	// access to struct's fields via pointers
	p.X = 1e9  // without dereference
	(*p).Y = 4 // with dereference
	fmt.Println("v", v)
}
