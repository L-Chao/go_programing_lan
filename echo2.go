package main

import (
	"fmt"
	"os"
)

func main()  {
	s, sep := "", ""
	//rang 产生一对值：索引、该索引处的元素值
	// _ 空标识符
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
