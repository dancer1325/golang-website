//go:build norun || OMIT

package main

import "fmt"

type I interface {
	M()
}

func main() {
	// nil interface
	var i I
	describe(i) // either value or type are nil
	i.M()       // NP in runtime
}

func describe(i I) {
	//  %v	-- value
	//  %T	-- type
	fmt.Printf("(%v, %T)\n", i, i)
}
