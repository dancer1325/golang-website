//go:build nobuild || OMIT

package main

import "golang.org/x/tour/pic"

type Image struct{}

// TODO:
func main() {
	m := Image{}
	pic.ShowImage(m)
}
