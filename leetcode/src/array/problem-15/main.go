package main

import (
	"sort"
)

func threeSum(numbers []int) [][]int {
	result := make([][]int, 0)
	if numbers == nil || len(numbers) < 3 {
		return result
	}

	length := len(numbers)

	sort.Ints(numbers)

	for first := 0; first < length; first++ {
		if first > 0 && numbers[first] == numbers[first-1] {
			continue
		}

		third := length - 1
		target := -1 * numbers[first]

		for second := first+1; second < length; second++ {
			if second > first + 1 && numbers[second] == numbers[second-1] {
				continue
			}

			for second < third && numbers[second] + numbers[third] > target {
				third--
			}

			if second == third {
				break
			}

			if numbers[second] + numbers[third] == target {
				result = append(result, []int{numbers[first], numbers[second], numbers[third]})
			}
		}
	}

	return result
}