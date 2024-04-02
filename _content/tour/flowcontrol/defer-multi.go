//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	// NOT deferred
	fmt.Println("counting")

	// defer happens in stack == last-in-first-out
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	// function deferred invoked here previously
	fmt.Println("done")
}
