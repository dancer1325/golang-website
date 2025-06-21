//go:build OMIT

package main

import "fmt"

func main() {
	// slice literal
	s := []int{2, 3, 5, 7, 11, 13}

	// reassign s, taking another size
	s = s[1:4] // 0 is the first element
	fmt.Println(s)

	// reassign s, taking another size
	// lowestIndex NOT specified -> by default it's 0
	s = s[:2] // == s[0:2]
	fmt.Println(s)

	// reassign s, taking another size
	// highestIndex NOT specified -> by default its slice's length
	s = s[1:] // == s[1:2]
	fmt.Println(s)
}
