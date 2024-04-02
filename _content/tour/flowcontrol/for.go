//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	sum := 0 // short-declaration variable

	// ()		NOT valid
	//for (i := 0; i < 10; i++) {
	// for initStatement; condition; postStatement {...}
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
