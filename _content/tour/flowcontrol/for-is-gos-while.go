//go:build OMIT

package main

import "fmt"

func main() {
	sum := 1

	// for condition {...}
	// == other program language's while
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
