//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	// array
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(fmt.Sprintf("primes type %T", primes))

	// slices
	var s []int = primes[1:4] // 0 is the first element		& highestIndex is NOT included
	fmt.Println(fmt.Sprintf("s type %T", s), s)
}
