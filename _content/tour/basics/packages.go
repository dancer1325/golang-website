//go:build OMIT

package main

import (
	"fmt"
	"math/rand" // == all files containing package rand -- Go into it, to check
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
