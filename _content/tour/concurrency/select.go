//go:build OMIT

package main

import "fmt"

// arguments	-- several channels of type int
func fibonacci(c, quit chan int) {
	x, y := 0, 1

	// for { select { case1: .. case2: .. }}
	for {
		select {
		case c <- x: // send x to the channel c
			x, y = y, x+y
		case <-quit: // quit channel receives a value
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// create a channel
	c := make(chan int)

	// create a channel
	quit := make(chan int)

	// goroutine -- creating the function to execute directly here --
	go func() {
		// for initStatement; condition; postStatement {...}
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // print the values sent by the channel c
		}
		quit <- 0 // send 0 to the channel quit
	}()

	fibonacci(c, quit) // same goroutine that main one
}
