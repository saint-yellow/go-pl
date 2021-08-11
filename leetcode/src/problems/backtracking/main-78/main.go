// LeetCode 主站 Problem Nr. 78: 子集

package main

import (
	"fmt"
)

func subsets(numbers []int) [][]int {
	return method1(numbers)
}

func method1(numbers []int) [][]int {
	result := make([][]int, 0)
	track := make([]int, 0)

	var backtracking func(numbers []int, startIndex int)
	backtracking = func(numbers []int, startIndex int) {
		temp := make([]int, len(track))
		copy(temp, track)
		result = append(result, temp)

		if startIndex >= len(numbers) {
			return
		}

		for i := startIndex; i < len(numbers); i++ {
			track = append(track, numbers[i])
			backtracking(numbers, i+1)
			track = track[:len(track)-1]
		}
	}

	backtracking(numbers, 0)
	return result
}

func main() {
	numbers := []int{1,2,3,4}
	fmt.Println(subsets(numbers))
}