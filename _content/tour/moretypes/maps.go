//go:build OMIT

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map[keyType]valueType
var m map[string]Vertex

func main() {
	// make(map[keyType]valueType)
	m = make(map[string]Vertex)
	fmt.Println(m) // empty initialized map
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
