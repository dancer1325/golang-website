//go:build OMIT
// +build OMIT

package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// method with pointer as receiver -> accepts either value or pointer!!
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// function with pointer as argument -> JUST pointers are valid!!
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2) // method with pointer as argument -> type are ALSO valid!!
	fmt.Println("Type defined - Via pointer receiver - after scale", v)
	//ScaleFunc(v, 10)				-- Function with pointer arguments --> MUST pass pointers
	ScaleFunc(&v, 10)
	fmt.Println("Type defined - Via pointer function - after scale the pointer", v)

	p := &Vertex{3, 4}
	p.Scale(2)
	fmt.Println("Pointer defined - Via pointer receiver - after scale ", p)
	ScaleFunc(p, 10)
	fmt.Println("Pointer defined - Via pointer function - after scale ", p)
}
