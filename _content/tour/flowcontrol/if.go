//go:build OMIT

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// ()			NOT valid -- in fact, GoLand IDE removes it automatically
	//if (x < 0) {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
