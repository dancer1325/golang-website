//go:build OMIT

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map[keyType]valueType{…} 			— map literals —
var m = map[string]Vertex{
	// ️if the top level type is type name → you can omit it		-- compare with map-literals.go --
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
