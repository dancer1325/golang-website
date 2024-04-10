//go:build OMIT
// +build OMIT

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// Method on type *T  -> variable of type *T can be assigned to type I
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T

	// With interface's value nil
	i = t
	describe(i)
	i.M() // NO null pointer is triggered

	// type literal
	i = &T{"hello"}
	describe(i)
	i.M() // interface's value exist here
}

func describe(i I) {
	//  %v	-- value
	//  %T	-- type
	fmt.Printf("(%v, %T)\n", i, i)
}
