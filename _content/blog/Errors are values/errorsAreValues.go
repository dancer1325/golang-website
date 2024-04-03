package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

}

func abstractInASeparatedMethod() {
	// Open the file
	file, err := os.Open("example.txt")
	fmt.Print("err ", err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		token := scanner.Text()
		fmt.Print("token ", token)
	}
	if err := scanner.Err(); err != nil {
		// process the error
	}
}

// TODO:
