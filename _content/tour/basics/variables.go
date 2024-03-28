//go:build OMIT
// +build OMIT

package main

import "fmt"

var c, python, java bool

// TODO: Check `byte` alias for `uint8`
// TODO: Check `rune` alias for `int32` & Unicode code point

// factor into blocks
var (
	another bool
	yes     bool   = true
	name    string = "Noelia"
)

func main() {
	var i int
	fmt.Println(i, c, python, java, another, name)
}
