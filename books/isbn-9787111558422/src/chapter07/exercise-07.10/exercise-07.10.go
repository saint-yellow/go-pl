package main

import (
	"fmt"
	"sort"
)

func main() {
	s1, s2 := StringSort("ada"), StringSort("asd")
	fmt.Println(IsPalindrome(s1), IsPalindrome(s2))

	r1, r2 := RuneSort([]rune("ada")), RuneSort([]rune("asd"))
	fmt.Println(IsPalindrome(r1), IsPalindrome(r2))
}

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i != j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		} else {
			return false
		}
	}
	return true
}

type StringSort string
func (s StringSort) Len() int {return len(s)}
func (s StringSort) Less(i, j int) bool {return s[i] < s[j]}
func (s StringSort) Swap(i, j int) {
	// Golang's string type is immutable.
}

type RuneSort []rune
func (s RuneSort) Len() int {return len(s)}
func (s RuneSort) Less(i, j int) bool {return s[i] < s[j]}
func (s RuneSort) Swap(i, j int) {s[i], s[j] = s[j], s[i]}
