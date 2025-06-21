//go:build OMIT

package main

// function functionName[typeParameters restrictionsToTypeParameters](arguments) returnedType {…}
/*
T is a type parameter
returns the index of x in s, else -1
*/
func Index[T comparable](s []T, x T) int {
	// for first, second range sliceOrMap {…}
	// if slice -> first == slice item's position & second == copy of the slice item’s value
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	// It can be applied to any type implementing the comparable interface
}
