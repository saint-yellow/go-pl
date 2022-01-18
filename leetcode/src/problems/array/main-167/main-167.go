package main

func twoSum(numbers []int, target int) []int {
	return method1(numbers, target)
}

func method1(numbers []int, target int) []int {
	result := make([]int, 2)
	left, right := 0, len(numbers)-1
	for left <= right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			result[0], result[1] = left+1, right+1
			break
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return result
}
