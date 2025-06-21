//go:build OMIT

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Same as 'methods-pointers' but as function instead of method
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Same as 'methods-pointers' but as function instead of method
func Scale(v *Vertex, f float64) {
	// pointer -> real argument passed is modified
	v.X = v.X * f
	v.Y = v.Y * f
}

// Same as 'methods-pointers' but as function instead of method
func ScaleAsType(v Vertex, f float64) {
	// v here is a copy ONCE it's invoked this function!!!
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println("Via pointer function - after scale ", v)
	fmt.Println("Via pointer function - after Abs ", Abs(v))

	w := Vertex{3, 4}
	ScaleAsType(w, 10)
	fmt.Println("Via type function - after scale ", w)
	fmt.Println("Via type function - after Abs ", Abs(w))
}
