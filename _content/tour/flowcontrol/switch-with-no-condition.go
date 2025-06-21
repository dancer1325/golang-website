//go:build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// switch initStatement; condition { case....}
	// initStatement			Optional
	// condition				Optional -- if it's missing -> always enters by the first case --
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
