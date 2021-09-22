// LeetCode 主站 Problem Nr. 977: 有序数组的平方

package main

import (
	"fmt"
	"sort"
)

func sortedSquares(numbers []int) []int {
	return method1(numbers)
}

func method1(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		numbers[i] = numbers[i] * numbers[i]
	}

	sort.Ints(numbers)

	return numbers
}

func method2(numbers []int) []int {
	result := make([]int, len(numbers))
	k := len(numbers) - 1

	for i, j := 0, len(numbers) - 1; i <= j; {
		if numbers[i] *numbers[i] < numbers[j] * numbers[j] {
			result[k] = numbers[j] * numbers[j]
			k--
			j--
		} else {
			result[k] = numbers[i] * numbers[i]
			k--
			i++
		}
	}

	return result
}

func main() {
	numbers := []int{-4, -1, 0, 3, 10}
	// fmt.Println(sortedSquares(numbers))
	fmt.Println(method2(numbers))
}