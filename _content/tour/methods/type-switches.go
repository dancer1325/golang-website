//go:build OMIT
// +build OMIT

package main

import "fmt"

func do(i interface{}) {
	// switch initStatement := interfaceName.(type) { case .... }
	switch v := i.(type) { // type is a keyword
	case int: // types are specified here
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true) // default case == v is of type bool
}
