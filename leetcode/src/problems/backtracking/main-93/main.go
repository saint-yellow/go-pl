// LeetCode 主站 Problem Nr. 93: 复原IP地址

package main

import (
	"fmt"
	"strconv"
	"strings"
)



func restoreIpAddresses(s string) []string {
	return method1(s)
}

func method1(s string) []string {
	result, path := make([]string, 0), make([]string, 0)

	isValidIP := func(s string, startIndex, end int) bool {
		checkInt, _ := strconv.Atoi(s[startIndex:end+1])
		if end - startIndex + 1 > 1 && s[startIndex] == '0' {
			return false
		}
		if checkInt > 255 {
			return false
		}

		return true
	}

	var backtracking func(s string, path []string, startIndex int, result *[]string)
	backtracking = func(s string, path []string, startIndex int, result *[]string) {
		if startIndex == len(s) && len(path) == 4 {
			temp := strings.Join(path, ".")
			*result = append(*result, temp)
		}

		for i := startIndex; i < len(s); i++ {
			path = append(path, s[startIndex:i+1])
			if i - startIndex + 1 <= 3 && len(path) <= 4 && isValidIP(s, startIndex, i) {
				backtracking(s, path, i+1, result)
			} else {
				return
			}

			path = path[:len(path)-1]
		}
	}

	if s == "" || len(s) > 12 {
		return result
	}

	backtracking(s, path, 0, &result)
	return result
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}