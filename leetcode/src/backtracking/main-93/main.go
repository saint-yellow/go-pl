// LeetCode 主站 Problem Nr. 93: 复原IP地址

package main

import (
	"fmt"
	"strings"
)

// func restoreIpAddresses(s string) []string {
// 	return method1(s)
// }

// func method1(s string) []string {
// 	result := make([]string, 0)

// 	isValid := func(s string, start, end int) bool {
// 		if start > end {
// 			return false
// 		}

// 		if s[start] == '0' && start != end {
// 			return false
// 		}

// 		number := 0
// 		for i := start; i <= end; i++ {
// 			if s[i] > '9' || s[i] < '0' {
// 				return false
// 			}

// 			number = number * 10 + int(s[i] - '0')
// 			if number > 255 {
// 				return false
// 			}
// 		}

// 		return true
// 	}

// 	var backtracking func(s string, startIndex, dotNumber int)
// 	backtracking = func(s string, startIndex, dotNumber int) {
// 		if dotNumber == 3 {
// 			if isValid(s, startIndex, len(s)-1) {
// 				result = append(result, s)
// 			}
// 			return
// 		}

// 		for i := startIndex; i < len(s); i++ {
// 			if isValid(s, startIndex, i) {

// 			}
// 		}
// 	}
// }


func main() {
	s := strings.Split("saintyellow", "")
	s = append(s, "")
	copy(s[6:], s[5:11])
	s[5] = "-"
	fmt.Println(s)
}