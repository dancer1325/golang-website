package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	interfaceTypeVariableStoreAPair()
}

/*
* variable of type interface stores a pair -- interface values (interfaceValue, type) --
 */
func interfaceTypeVariableStoreAPair() {
	var r io.Reader
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		return
	}

	// r		contains  the (value, type) pair --  (tty, *os.File)
	// can be assigned to a variable io.Reader, because	*os.File implements the method Read
	r = tty

	// *os.File implements methods other than Read -> we can do
	var w io.Writer
	w = r.(io.Writer) // type assertion
	fmt.Print(fmt.Sprintf("w type is %T", w))
}
