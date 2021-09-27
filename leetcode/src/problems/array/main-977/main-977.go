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
		a := numbers[i] * numbers[i]
		b := numbers[j] * numbers[j]
		if a < b {
			result[k] = b
			k--
			j--
		} else {
			result[k] = a
			k--
			i++
		}
	}

	return result
}

func main() {
	numbers := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(numbers))
	fmt.Println(method2(numbers))
}