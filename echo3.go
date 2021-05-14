package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args)
	for i, s := range os.Args {
		fmt.Print(i)
		fmt.Println(" " + s)
	}
}
