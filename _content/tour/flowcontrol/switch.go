//go:build OMIT
// +build OMIT

package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	fmt.Print("Go runs on ")
	const Darwin = "darwin"

	// switch initStatement; condition { case....}
	switch os := runtime.GOOS; os {
	// switch case can be a constant
	case Darwin:
		fmt.Println("OS X.")
		// break NOT needed
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	const favouriteNumber int = 4
	// initStatement		optional
	switch favouriteNumber {
	// switch case can be an integer
	case 2:
		fmt.Println("2")
	case 5:
		fmt.Println("5")
	default:
		fmt.Printf(strconv.Itoa(favouriteNumber))
	}

	// variables defined into switch's scope  -> NOT available outside  		-- Uncomment next line
	//fmt.Println(os)
}
