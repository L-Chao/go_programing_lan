package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}
func (s *IntSet) Has(x int) bool {
	word, bit := x / 64, uint(x % 64)
	return word < len(s.words) && (s.words[word] & (1 << bit) != 0)
}
func (s *IntSet) Add(x int) {
	word, bit := x / 64, uint(x % 64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] |= tWord
		} else {
			s.words = append(s.words, tWord)
		}
	}
}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word & (1 << uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64 * i + j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
func (s *IntSet) Len() int {
	ret := 0
	for _, word := range s.words {
		for word > 0 {
			ret += 1
			word &= word - 1
		}
	}
	return ret
}
func (s *IntSet) Remove(x int) {
	word, bit := x / 64, uint(x % 64)
	s.words[word] &^= 1 << bit
}
func (s *IntSet) Clear() {
	s.words = []uint64{}
}
func (s *IntSet) Copy() *IntSet {
	ret := new(IntSet)
	for _, word := range s.words {
		ret.words = append(ret.words, word)
	}
	return ret
}
func main() {
	var x IntSet
	//y := IntSet{}
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println("x = ", x.String())
	y := x.Copy()
	fmt.Println("y = ", y.String())

	fmt.Println("x.Len() = ", x.Len())
	x.Remove(1)
	fmt.Println(x.String())
	x.Clear()
	fmt.Println(x.String())
	/*
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123))

	 */
}
