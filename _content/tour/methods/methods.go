//go:build OMIT

package main

import (
	"fmt"
	"math"
)

// NOT class exist

type Vertex struct {
	X, Y float64
}

// Method
// (v Vertex)		receiver -- in this case of type Vertex --
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
