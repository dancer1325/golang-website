//go:build OMIT

package main

import "fmt"

// "factor"  `()` into blocks
const (
	Pi     = 3.14
	Random = 2
)

func main() {
	const World = "世界"
	fmt.Println(World, Pi, "Day")
	fmt.Println(fmt.Sprintf("World constant is of type %T and Pi constant is of type %T", World, Pi))

	const Truth = true
	fmt.Println(Truth)
	fmt.Println(fmt.Sprintf("Truth constant is of type %T", Truth))

	// NOT valid to declare via      `:`    -- Uncomment next line to check --
	//const invalid : 3

	//TODO: How to declare a constant of type character?
	const myRune rune = 'A'
	fmt.Println(fmt.Sprintf("myRune is of type %T", myRune))
}
