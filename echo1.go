package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// := short variable declaration
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}