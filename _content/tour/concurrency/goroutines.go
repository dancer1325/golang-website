//go:build OMIT

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	// for initStatement; condition; postStatement {...}
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// go function(argumentsToPass)
	go say("world")

	// Runs in the current goroutine -> executed before, without waiting the finish of other new goroutine
	say("hello")
}
