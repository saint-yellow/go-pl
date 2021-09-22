// LeetCode 主站 Problem Nr. 27: 移除元素

package main

func removeElement(numbers []int, value int) int {
	return method1(numbers, value)
}

func method1(numbers []int, value int) int {
	size := len(numbers)
	for i := 0; i < size; i++ {
		if numbers[i] == value {
			for j := i + 1; j < size; j++ {
				numbers[j-1] = numbers[j]
			}
			i--
			size--
		}
	}
	return size
}

func method2(numbers []int, value int) int {
	slow := 0
	for fast := 0; fast < len(numbers); fast++ {
		if value != numbers[fast] {
			slow++
			numbers[slow] = numbers[fast]
		}
	}
	return slow
}