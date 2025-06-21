//go:build OMIT

package main

import "fmt"

// type structName struct { fields }
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
