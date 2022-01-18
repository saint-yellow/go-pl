package main

import (
	"fmt"
	"sort"
)

func threeSum(numbers []int) [][]int {
	return method1(numbers)
}


func method1(numbers []int) [][]int {
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

func method2(numbers []int) [][]int {
	sort.Ints(numbers)

	result := [][]int{}
	for i := 0; i < len(numbers)-2; i++ {
		n1 := numbers[i]
		if n1 > 0 {
			break
		}

		if i > 0 && n1 == numbers[i-1] {
			continue
		}

		left, right := i+1, len(numbers)-1
		for left < right {
			n2, n3 := numbers[left], numbers[right]

			if n1 + n2 + n3 == 0 {
				result = append(result, []int{n1, n2, n3})

				for left < right && numbers[left] == n2 {
					left++
				}

				for left < right && numbers[right] == n3 {
					right--
				}
			} else if n1 + n2 + n3 > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, 4}))
	fmt.Println(method2([]int{-1, 0, 1, 2, -1, 4}))
}