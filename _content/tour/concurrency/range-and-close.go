//go:build OMIT

package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	// short declaration
	x, y := 0, 1

	// for initStatement; condition; postStatement {...}
	for i := 0; i < n; i++ {
		c <- x // send values to the channel
		x, y = y, x+y
	}

	// close the channel  -- Reason: receiver goroutine has got a for .. range  --
	close(c)
}

func main() {
	// Create a buffered channel
	c := make(chan int, 10)

	// goroutine
	go fibonacci(cap(c), c)

	// for variableDeclaration := range channelVariable {…}
	for i := range c { // infinite loop till channel is closed!!
		fmt.Println(i)
	}
}
