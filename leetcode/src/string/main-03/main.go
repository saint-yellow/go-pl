package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	return method1(s)
}

func method1(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}

	length := 1
	left , right := 0, 1
	
	for right < n {
		if !contains(s[left:right], s[right]) {
			right += 1
		} else {
			length = max(length, right-left)
			left += 1
		}
	}
	length = max(length, n-left)
	return length
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func contains(s string, b byte) bool {
	for _, v := range s {
		if v == rune(b) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(method1("abcabcbb"))
	fmt.Println(method1("bbbbb"))
	fmt.Println(method1("pwwkew"))
}