package main

func search(numbers []int, target int) int {
	return method2(numbers, target)
}

func method1(numbers []int, target int) int {
	left, right := 0, len(numbers)-1

	for left <= right {
		middle := (left+right)/2

		if target == numbers[middle] {
			return middle
		} else if target < numbers[middle] {
			right = middle-1
		} else {
			left = middle+1
		}
	}

	return -1
}

func method2(numbers []int, target int) int {
	left, right := 0, len(numbers)

	for left < right {
		middle := (left+right)/2

		if target == numbers[middle] {
			return middle
		} else if target < numbers[middle] {
			right = middle
		} else {
			left = middle+1
		}
	}

	return -1
}