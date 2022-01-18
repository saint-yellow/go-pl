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
		if _, ok := mapping[i]; ok {
			mapping[i] += 1
		} else {
			mapping[i] = 1
		}
	}

	for _, i := range t {
		if v, ok := mapping[i]; ok && v > 0 {
			mapping[i] -= 1
		} else {
			return false
		}
	}

	return true
}

func method2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapping := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if v, ok := mapping[s[i]]; ok && v >= 0 {
			mapping[s[i]]++
		} else {
			mapping[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if v, ok := mapping[s[i]]; ok && v >= 1 {
			mapping[s[i]]--
		} else {
			return false
		}
	}

	return true
}

func method3(s, t string) bool {
	record := [26]int{}
	for _, v := range s {
		record[v-rune('a')]++
	}
	for _, v := range t {
		record[v-rune('a')]--
	}
	return record == [26]int{}
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}