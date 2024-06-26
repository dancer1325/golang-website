//go:build ignore && OMIT
// +build ignore,OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	boring("boring!")
}

// STOP OMIT

// less random deterministic
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
