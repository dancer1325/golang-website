//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	// declare a slice -- NO initialization --
	var s []int
	fmt.Println(s, len(s), cap(s))

	// zero value of a slice is nil
	if s == nil {
		fmt.Println("nil!")
	}
}
