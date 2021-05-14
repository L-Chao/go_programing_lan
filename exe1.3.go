package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main()  {
	loopS := time.Now()
	loop()
	costLoop := time.Since(loopS)
	fmt.Printf("loop spend : %d\n", costLoop)
	noLoopS := time.Now()
	noLoop()
	costNoLoop := time.Since(noLoopS)
	fmt.Printf("no loop spend : %d\n", costNoLoop)
}

func loop() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func noLoop() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}