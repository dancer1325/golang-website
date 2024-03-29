//go:build ignore || OMIT
// +build ignore OMIT

package main

import (
	"fmt"
	"math"
)

/*
* interface type
 */
type Abser interface {
	// set of method signatures
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// At runtime, the type  is changed
	fmt.Println(fmt.Sprintf("Previous to assign to f %T", a))
	a = f // MyFloat has an implementation of Abs() method -> MyFloat implements Abser
	fmt.Println(fmt.Sprintf("After assigning to f %T", a))
	a = &v // *Vertex has an implementation of Abs() method -> *Vertex implements Abser
	fmt.Println(fmt.Sprintf("After assigning to v %T", a))

	// In the following line, v is a Vertex (NOT *Vertex)
	// Vertex has NOT implementation of Abs() method -> Vertex NOT implement Abser -> Error in runtime (Uncomment next line to check)
	//a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

// Method on type MyFloat
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// Method on type *Vertex      NOT -> Method on type Vertex
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
