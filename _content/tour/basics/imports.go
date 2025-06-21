//go:build OMIT

package main

// ()			-- 'factor' allows factoring all the imports
import (
	"fmt"
	"math"
	//"aa  /here"		// path can NOT contain whitespace
)

func main() {
	// %g		float points in compact representation
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
