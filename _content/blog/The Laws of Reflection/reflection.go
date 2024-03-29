package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

type MyIntAnother int

func main() {
	staticallyTyped()
	interfaceTypeDefinitionWithReadAsExample() // NO error in runtime == all good
}

/**
* although several types have the same underlying type -> they are distinct
 */
func staticallyTyped() {
	var i int
	var j MyIntAnother

	// Statically typed -- Uncomment next line
	//i==j

	// Dynamically (== Runtime) they are also distinct types
	if reflect.TypeOf(i) == reflect.TypeOf(j) {
		fmt.Println("i and j have the same type")
	} else {
		fmt.Println("i and j do NOT have the same type")
	}
}

/*
* Example to clarify interface type definition
* A variable of type io.Reader can hold any value / has a `Read` method -- 1! Reader's method --
 */
func interfaceTypeDefinitionWithReadAsExample() {
	var r io.Reader
	r = os.Stdin
	r = bufio.NewReader(r)
	r = new(bytes.Buffer)
}
