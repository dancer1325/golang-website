//go:build nobuild || OMIT

package main

import (
	"fmt"
	"math"
)

func main() {
	// pi		NOT exported names -> can NOT be referred -- Uncomment next line --
	//fmt.Println(math.pi)
	fmt.Println(math.Pi)
}
