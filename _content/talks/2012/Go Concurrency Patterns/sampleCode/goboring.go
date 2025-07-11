//go:build ignore && OMIT
// +build ignore,OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Run the boring() as go routine
	go boring("boring!") // HL
}

// STOP OMIT

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
