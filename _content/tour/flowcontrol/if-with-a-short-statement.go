//go:build OMIT

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// if initStatement; condition {…}
	if v := math.Pow(x, n); v < lim {
		return v
	}

	// variables defined in the initStatement are visible just into if scope
	//fmt.Println(v)

	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
