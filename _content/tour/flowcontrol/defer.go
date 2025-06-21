//go:build OMIT

package main

import "fmt"

func main() {
	// next function is deferred
	defer fmt.Println("world")

	// This function is executed first
	fmt.Println("hello")
}
