//go:build OMIT
// +build OMIT

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Method of a receiver as type
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Method of a receiver as pointer
func (v *Vertex) Scale(f float64) {
	// pointer -> real argument passed is modified
	v.X = v.X * f
	v.Y = v.Y * f
}

// Same method as before but as type
func (v Vertex) ScaleAsType(f float64) {
	// v here is a copy ONCE it's invoked this method!!!
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println("Via pointer receiver - after scale ", v)
	fmt.Println("Via pointer receiver - after Abs ", v.Abs())

	w := Vertex{3, 4}
	w.ScaleAsType(10)
	fmt.Println("Via type receiver - after scale ", w)
	fmt.Println("Via type receiver - after Abs ", w.Abs())
}
