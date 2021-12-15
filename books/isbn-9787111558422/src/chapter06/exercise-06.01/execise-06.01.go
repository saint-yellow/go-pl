package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buffer bytes.Buffer
	buffer.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buffer.Len() > len("{") {
					buffer.WriteByte(' ')
				}
				fmt.Fprintf(&buffer, "%d", 64*i+j)
			}
		}
	}
	buffer.WriteByte('}')
	return buffer.String()
}

func (s *IntSet) Len() int {
	return 0
}

func (s *IntSet) Remove(x int) {

}

func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)
}

func (s *IntSet) Copy() (r *IntSet) {
	var t IntSet
	t.words = append(t.words, s.words...)
	return &t
}

func main() {
	s := IntSet{}
	t := s.Copy()
	fmt.Printf("%x\n", &s)
	fmt.Printf("%x\n", &t)
}