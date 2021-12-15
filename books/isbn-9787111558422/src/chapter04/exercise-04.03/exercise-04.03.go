package main

import "fmt"

func reverse(s *[9]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(s)

	reverse(&s)
	fmt.Println(s)
}
