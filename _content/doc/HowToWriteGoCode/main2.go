package main

// TODO: Pending to study the notes about it
import (
	"fmt"                          // prefix path NOT existing because it comes from standard library
	"github.com/google/go-cmp/cmp" // github.com/google/go-cmp  -- module path --  & /cmp	is the subditorectory within module
	//"golang.org/x/website/_content/doc/HowtoWriteGoCode/morestrings"
)

func main() {
	//fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	// TODO: How to access variable declared in another file?
	//fmt.Println(visibleRestPackageFiles)
	// also valid defining with capital as first letter
	//fmt.Println(VisibleRestPackageFiles)

	// TODO: How to access functions declared in another file?
	//ReverseRunes("Alfred")
}
