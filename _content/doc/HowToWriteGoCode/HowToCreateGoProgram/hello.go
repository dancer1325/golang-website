// First statement of the Go program
//package HowToCreateGoProgram

// To have executable commands
package main

import (
	"example/user/hello/morestrings" // ModName/PackageName
	"fmt"
)

func main() {
	fmt.Println("Hello, world.")

	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}
