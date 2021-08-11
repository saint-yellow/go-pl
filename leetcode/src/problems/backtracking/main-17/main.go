// LeetCode Problem Nr. 17: 电话号码的字母组合

package main

import (
	"fmt"
	"strings"
)

var letterMap = map[byte]string {
	'0': "",
	'1': "",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}



func letterCombinations(digits string) []string {
	return method1(digits)
}

func method1(digits string) []string {
	result := make([]string, 0)
	if digits == "" {
		return result
	}

	track := make([]string, 0)

	var backtracking func(digitsCount, index int)
	backtracking = func(digitsCount, index int) {
		if index == digitsCount {
			result = append(result, strings.Join(track, ""))
			return
		}

		digit := digits[index]
		letters := letterMap[digit]
		for i := 0; i < len(letters); i++ {
			track = append(track, string(letters[i]))
			backtracking(digitsCount, index + 1)
			track = track[:len(track)-1]
		}
	}

	digitsCount := len(digits)
	backtracking(digitsCount, 0)
	return result
}

func main() {
	fmt.Println(letterCombinations("02"))
}