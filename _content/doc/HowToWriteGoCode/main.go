package main

// TODO: Pending to study the notes about it
import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	//"golang.org/x/website/_content/doc/HowtoWriteGoCode/morestrings"
)

func main() {
	//fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
