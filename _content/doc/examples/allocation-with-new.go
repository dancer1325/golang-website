package main

import (
	"bytes"
	"fmt"
	"sync"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func main() {
	// 1. allocates 0's
	p := new(SyncedBuffer) // type *SyncedBuffer 	== pointer
	fmt.Printf("p's type %T p's value %p\n", p, p)

	var v SyncedBuffer // type  SyncedBuffer
	fmt.Printf("v's type %T v's value %v\n", v, v)

	// 2. 0 works transitively		-- ALTHOUGH, it's 0 --
	p.lock.Lock() // EACH type's 0 value can be used -- WITHOUT -- further initialization
	p.buffer.WriteString("Hello from zero value!")
	p.lock.Unlock()
	fmt.Println("Content:", p.buffer.String())
	fmt.Printf("Type: %T, Value: %+v\n", p, p)
}
