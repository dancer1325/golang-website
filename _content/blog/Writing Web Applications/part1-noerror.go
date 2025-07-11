// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

/*
_	error NOT handled
*/
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename)
	return &Page{Title: title, Body: body}
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page.")}
	p1.save()
	p2 := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
