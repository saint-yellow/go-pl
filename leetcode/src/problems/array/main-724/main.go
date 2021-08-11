package main


func pivotIndex(numbers []int) int {
	totalSum := 0
	for _, n := range numbers {
		totalSum += n
	}
	leftSum := 0
	for i, n := range numbers {
		if n + 2 * leftSum == totalSum {
			return i
		}
		leftSum += n
	}
	return -1
}