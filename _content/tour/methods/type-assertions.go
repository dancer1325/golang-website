//go:build ignore || OMIT
// +build ignore OMIT

package main

import "fmt"

func main() {
	var i interface{} = "hello" // empty interface in this case

	// := interfaceName.(underlyingInterfaceValueType)
	// passing to 1 variable     -- interfaceValue’s value --
	s := i.(string)
	fmt.Println(s)

	// := interfaceName.(underlyingInterfaceValueType)
	// passing to 2 variables     -- interfaceValue’s value, boolean --
	s, ok := i.(string)
	fmt.Println(s, ok)

	// := interfaceName.(underlyingInterfaceValueType)			-- here passing wrong one --
	// passing to 2 variables     -- interfaceValue’s value, false --
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// := interfaceName.(underlyingInterfaceValueType)			-- here passing wrong one --
	// passing to 1 variable1     -- panic occurs --
	f = i.(float64) // panic
	fmt.Println(f)
}
