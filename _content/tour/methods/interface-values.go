//go:build OMIT

package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

// Method on type *T  -> variable of type *T can be assigned to type I
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

// Method on type F -> variable of type F can be assigned to type I
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"} // NO error, because *T implements I's method
	describe(i)
	i.M() // TODO: i is JUST interface's value? -> we are calling interface value's method

	i = F(math.Pi) // NO error, because F implements I's method
	describe(i)
	i.M()
}

/*
Display the interface values -- (value, type)
*/
func describe(i I) {
	//  %v	-- value
	//  %T	-- type
	fmt.Printf("(%v, %T)\n", i, i)
}
