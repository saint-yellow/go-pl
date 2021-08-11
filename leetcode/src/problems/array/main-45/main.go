package main

import "fmt"

func jump(numbers []int) int {
	return method1(numbers)
}

func method1(numbers []int) int {
	currentDistance, result, nextDistance := 0, 0, 0
	for i := 0; i < len(numbers)-1; i++ {
		nextDistance = max(numbers[i]+i, nextDistance)
		if i == currentDistance {
			currentDistance = nextDistance
			result++
		}
	}
	return result
}

func method2(numbers []int) int {
	result := 0

	// to-do: ...

	return result
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func main() {
	fmt.Println(jump([]int{2,3,1,1,4}))
	fmt.Println(jump([]int{2,3,0,1,4}))
}