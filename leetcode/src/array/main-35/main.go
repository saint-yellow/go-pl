package main

func searchInsert(numbers []int, target int) int {
	return method2(numbers, target)
}

func method1(numbers []int, target int) int {
	for index, value := range numbers {
		if target < value {
			return index
		}
	}

	return len(numbers)
}

func method2(numbers []int, target int) int {
	left, right := 0, len(numbers)-1
	for left <= right {
		middle := (left+right)/2
		if target < numbers[middle] {
			right = middle-1
		} else if target > numbers[middle] {
			left = middle+1
		} else {
			return middle
		}
	}
	return right+1
}

func method3(numbers []int, target int) int {
	left, right := 0, len(numbers)
	for left < right {
		middle := (left+right)/2
		if target < numbers[middle] {
			right = middle
		} else if target > numbers[middle] {
			left = middle+1
		} else {
			return middle
		}
	}
	return right+1
}