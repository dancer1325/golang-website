//go:build ignore && OMIT
// +build ignore,OMIT

package main

func main() {
	var value int

	// START1 OMIT
	// Declaring and initializing.
	var c chan int // channel of int
	c = make(chan int)
	// or
	c := make(chan int) // HL  -- Declare & initialize in 1! step
	// STOP1 OMIT

	// START2 OMIT
	// Sending on a channel.
	c <- 1 // HL
	// STOP2 OMIT

	// START3 OMIT
	// Receiving from a channel.
	// The "arrow" indicates the direction of data flow.
	value = <-c // HL
	// STOP3 OMIT

	_ = value
}
