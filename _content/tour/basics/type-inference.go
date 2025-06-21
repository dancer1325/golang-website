//go:build OMIT

package main

import "fmt"

func main() {
	v := 42 // change me!
	w := v  // type inferred also
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("w is of type %T\n", w)

	// Based on the precision
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}
