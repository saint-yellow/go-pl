// LeetCode 主站 Problem Nr. 209: 长度最小的子数组

package main

func minSubArrayLen(target int, numbers []int) int {
	return method1(target, numbers)
}

func method1(target int, numbers []int) int {
	i := 0
	l := len(numbers)
	sum := 0
	result := l + 1

	for j := 0; j < l; j++ {
		sum += numbers[j]
		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= numbers[i]
			i++
		}
	}

	if result == l + 1 {
		return 0
	} else {
		return result
	}
}