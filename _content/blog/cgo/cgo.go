package rand

/*
#include <stdlib.h>
*/
import "C" // pseudo-package -- interpreted by cgo as reference to -- C's name space

func Random() int {
	return int(C.random()) // C.random() returns C.long & -- via `int()` is converted to -- Go type
}

func Seed(i int) {
	C.srandom(C.uint(i)) // Go int -- via `C.uint() is converted to -- C unsigned int type
}
