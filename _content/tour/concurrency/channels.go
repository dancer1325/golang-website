//go:build OMIT
// +build OMIT

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0

	// for first, second range sliceOrMap {…}
	// first  _		because we want to skip it == NOT use it
	for _, v := range s {
		sum += v
	}

	fmt.Println("sum for s", s)

	c <- sum // send sum to c
}

func main() {
	// slice literal
	s := []int{7, 2, 8, -9, 4, 0}

	// initialize channel
	c := make(chan int)

	// goroutines using the same channel
	go sum(s[:len(s)/2], c) // s[:len(s)/2] == s[:3] == [7,2,8]
	go sum(s[len(s)/2:], c) // s[len(s)/2] == s[3] == [-9,4,0]

	fmt.Println("main")

	// This goroutine waits receiving the values from the other goroutines through the channel
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
