// LeetCode 主站 Problem Nr. 242: 有效的字幕异位词

package main

import "fmt"

func isAnagram(s, t string) bool {
	return method1(s, t)
}

func method1(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapping := make(map[rune]int)

	for _, i := range s {
		if v, ok := mapping[i]; ok {
			mapping[i] = v + 1
		} else {
			mapping[i] = 1
		}
	}

	for _, i := range t {
		if v, ok := mapping[i]; ok && v > 0 {
			mapping[i] = v - 1
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}