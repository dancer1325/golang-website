// First statement of the Go program
//package HowToCreateGoProgram

// To have executable commands
package main

import (
	"example/user/hello/morestrings" // ModName/PackageName
	"fmt"
	"github.com/google/go-cmp/cmp" // Packages from remote modules
)

func main() {
	fmt.Println("Hello, world.")

	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
