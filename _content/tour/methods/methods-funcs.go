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

// == Abs() in methods.go but coded as function with special receiver argument
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
