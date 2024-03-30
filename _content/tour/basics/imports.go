//go:build OMIT
// +build OMIT

package main

// ()			-- 'factor' allows factoring all the imports
import (
	"fmt"
	"math"
)

func main() {
	// %g		float points in compact representation
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
