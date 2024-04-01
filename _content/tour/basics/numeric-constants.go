//go:build OMIT
// +build OMIT

package main

import "fmt"

// "factor"  `()` into blocks
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int {
	// infer numeric constant value
	var intNumber = x*10 + 1
	fmt.Println(fmt.Sprintf("intNumber is of type %T", intNumber))
	return intNumber
}
func needFloat(x float64) float64 {
	// infer numeric constant value
	var float = x * 0.1
	fmt.Println(fmt.Sprintf("float is of type %T", float))
	return float
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
