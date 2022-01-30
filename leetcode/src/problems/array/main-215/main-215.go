package main

import (
	"fmt"
	"sort"
)

func findKthLargest(numbers []int, k int) int {
	return method1(numbers, k)
}

func method1(numbers []int, k int) int {
	sort.Ints(numbers)
	fmt.Println(numbers)
	return numbers[len(numbers)-k]
}

func main() {
	numbers := []int{3,2,1,6,5,4}
	fmt.Println(findKthLargest(numbers, 2))
}