//go:build OMIT

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// Method on type *MyError, which returns the built-in Error() function -> *MyError implements the built-in type error!!
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// if initStatement; condition {...}
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
