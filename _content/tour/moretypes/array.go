//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	// array of 2 values of typ string
	var a [2]string
	a[0] = "Hello" // 0 is the first element
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	// can NOT be resized		-- Uncomment next line --
	//primes[9] = 111
	fmt.Println(primes)

	// array literal
	var arrayLiteral = [3]bool{true, true, false}
	fmt.Println(arrayLiteral)
}
