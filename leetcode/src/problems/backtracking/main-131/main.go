// LeetCode 主站 Problem Nr. 131: 分割回文串

package main

import "fmt"

func partition(s string) [][]string {
	return method1(s)
}

func method1(s string) [][]string {
	result := make([][]string, 0)
	if s == "" {
		result = append(result, []string{s})
		return result
	}

	isPalindrome := func(left, right int) bool {
		for i, j := left, right; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	track := make([]string, 0)

	var backtracking func(n, startIndex int)
	backtracking = func(n, startIndex int) {
		if startIndex >= n {
			temp := make([]string, len(track))
			copy(temp, track)
			result = append(result, temp)
			return
		}

		for i := startIndex; i < n; i++ {
			if isPalindrome(startIndex, i) {
				track = append(track, s[startIndex:i+1])
			} else {
				continue
			}

			backtracking(n, i + 1)
			
			track = track[:len(track)-1]
		}
	}

	backtracking(len(s), 0)
	return result
}

func main() {
	p := partition("aab")
	fmt.Println(p)
}