// LeetCode 主站 Problem Nr. 344: 反转字符串

package main

func reverseString(s []byte) {
	method1(s)
}

func method1(s []byte) {
	if s == nil {
		return
	}

	start, end := 0, len(s)-1
	for start <= end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}