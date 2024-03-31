//go:build OMIT
// +build OMIT

package main

import "fmt"

// x & y are int -> int for x declaration is omitted
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
