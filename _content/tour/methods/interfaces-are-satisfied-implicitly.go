//go:build OMIT

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// Method on type T -> 👀type T implements the interface I👀
// NOT need to explicitly declare that it does so
func (t T) M() {
	// body of the method can be delegated to another package
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
