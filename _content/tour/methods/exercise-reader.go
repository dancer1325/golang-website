//go:build nobuild || OMIT

package main

import (
	"fmt"
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (MyReader) Read(b []byte) (n int, err error) {
	fmt.Println("Valid MyReader.Read")
	return len(b), err
}

func main() {
	reader.Validate(MyReader{})
}
