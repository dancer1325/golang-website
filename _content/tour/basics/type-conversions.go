//go:build OMIT
// +build OMIT

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4

	// type conversions
	var f float64 = math.Sqrt(float64(x*x + y*y)) // float64()
	var z uint = uint(f)                          // uint()

	fmt.Println(x, y, z)

	// explicit conversion is required  -- uncomment next lines to check that it fails in runtime --
	//var i int = 42
	//var f float64 = i				// float64(i)		required
	//var u uint = f					// uint(f)			required

}
