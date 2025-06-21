//go:build OMIT

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method with value as receiver -> accepts either value or pointer!!
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// function with value as argument -> JUST values are valid!!
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("Type defined - Via function ", AbsFunc(v))
	//fmt.Println("Pointer defined - Via function ", AbsFunc(&v))		-- Function with value arguments --> MUST pass values

	p := &Vertex{4, 3}
	fmt.Println("Pointer defined - Via method ", p.Abs()) // Valid to pass pointers to a function with value arguments
	fmt.Println("Type defined - Via function", AbsFunc(*p))
}
