//go:build OMIT

package main

import "fmt"

// (returnType1, returnType2)
func swap(x, y string) (string, string) {
	return y, x
}

// (returnType1, returnType2, returnType3)		-- Any number of returnedTypes are legal
func swapWithMore(x, y string) (string, string, string) {
	return y, x, "Hello"
}

// returnType1, returnType2		is NOT possible
/*func swapWithoutParenthesis(x, y string) string, string {
	return y, x
}*/

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	c, d, e := swapWithMore("Alfred", "Rosi")
	fmt.Println(c, d, e)
}
