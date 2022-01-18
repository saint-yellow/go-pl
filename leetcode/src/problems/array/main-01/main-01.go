// LeetCode 主站 Problem Nr. 1: 两数之和

package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	return method1(numbers, target)
}

func method1(numbers []int, target int) []int {
	dict := make(map[int]int)
	for i, n := range numbers {
		a := target - n
		if _, ok := dict[n]; ok {
			return []int{dict[n], i}
		}
		dict[a] = i
	}

	return []int{-1, -1}
}

func method2(numbers []int, target int) []int {
	m := make(map[int]int)
	for index, value := range numbers {
		if preIndex, ok := m[target-value]; ok {
			return []int{preIndex, index}
		} else {
			m[value] = index
		}
	}
	return []int{-1, -1}
}

func method3(numbers []int, target int) []int {
	for i1 := range numbers {
		for i2 := i1 + 1; i2 < len(numbers); i2++ {
			if target == numbers[i1] + numbers[i2] {
				return []int{i1, i2}
			}
		} 
	}
	return []int{-1, -1}
}

func main() {
	numbers := []int{2,7,11,5}
	target := 9
	fmt.Println(twoSum(numbers, target))
}
