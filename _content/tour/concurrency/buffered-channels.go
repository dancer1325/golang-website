//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2 // If you comment this line and NOT previous -> the program fails, since ALL channel must be completed, previous to READ their values
	//ch <- 6   // If the channel -> you can NOT send more values to the channel
	fmt.Println(<-ch)
	//fmt.Println(<-ch)			// If you comment it this one, program still runs

	ch <- 3 // Since now, it's empty -> you can send values to the channel
}
