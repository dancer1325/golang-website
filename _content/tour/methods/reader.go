//go:build OMIT
// +build OMIT

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	// create a slice of byte
	b := make([]byte, 8)
	// condition omitted -> forever loop till we break
	for {
		n, err := r.Read(b) // strings 		implement Read
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
