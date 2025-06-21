//go:build OMIT

package main

import "fmt"

// func(int) int		is all the returned function value
func adder() func(int) int {
	sum := 0
	// sum is outside its body -- function closure --
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
