//go:build OMIT

package main

import "fmt"

func main() {
	// create a map
	m := make(map[string]int)

	// adjust or insert an item
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// adjust or insert an item
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// delete the item entered
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// value, booleanIfExistKey := mapVariable[keyToRetrieveValue]
	v, ok := m["Answer"] // short declaration
	fmt.Println("The value:", v, "Present?", ok)

	// value := mapVariable[keyToRetrieveValue]			-- booleanIfExistKey is optional --
	w := m["Answer"]
	fmt.Println("w ", w)
}
