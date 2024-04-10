//go:build ignore || OMIT
// +build ignore OMIT

package main

import (
	"fmt"
	"golang.org/x/tour/reader"
	"io"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13Reader) Read(b []byte) (n int, err error) {
	fmt.Println("Valid rot13Reader.Read")
	return len(b), err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	reader.Validate(r)
	//io.Copy(os.Stdout, &r)		// TODO: Fix because it's an infinite loop
}
