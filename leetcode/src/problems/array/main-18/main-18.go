// LeetCode 主站 Problem Nr. 18: 四数之和

package main

import (
	"fmt"
	"sort"
)

func fourSum(numbers []int, target int) [][]int {
	return method1(numbers, target)
}

func method1(numbers []int, target int) [][]int {
	length := len(numbers)

	if length < 4 {
		return nil
	}

	sort.Ints(numbers)
	var result [][]int
	for i := 0; i < length-3; i++ {
		n1 := numbers[i]

		if i > 0 && n1 == numbers[i-1] {
			continue
		}

		for j := i+1; j < length-2; j++ {
			n2 := numbers[j]

			if j > i+1 && n2 == numbers[j-1] {
				continue
			}

			left, right := j+1, length-1
			for left < right {
				n3, n4 := numbers[left], numbers[right]
				sum := n1 + n2 + n3 + n4

				if sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					result = append(result, []int{n1, n2, n3, n4})

					for left < right && n3 == numbers[left+1] {
						left++
					}

					for left < right && n4 == numbers[right-1] {
						right--
					}

					right--
					left++
				}
			}
		}
	}


	return result
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2}, 8))
}