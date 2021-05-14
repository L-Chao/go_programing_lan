package main

import (
	"bytes"
	"fmt"
	"sort"
)

type StringSlice []string
func (p StringSlice) Len() int {
	return len(p)
}
func (p StringSlice) Less(i int, j int) bool {
	return p[i] < p[j]
}
func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p StringSlice) String() string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, str := range p {
		buf.WriteString(str)
		if i < p.Len() - 1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']')
	return buf.String()
}
func main () {
	s := StringSlice{"a", "c", "s", "d", "b"}
	sort.Sort(s)
	fmt.Println(s.String())
}
