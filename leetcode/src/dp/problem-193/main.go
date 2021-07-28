package main

import "fmt"

func rob(numbers []int) int {
	return method2(numbers)
}

func method1(numbers []int) int {
	n := len(numbers)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return numbers[0]
	}

	dp := make([]int, n)
	dp[0] = numbers[0]
	dp[1] = max(numbers[0], numbers[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+numbers[i])
	}
	return dp[n-1]
}

func method2(numbers []int) int {
	n := len(numbers)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return numbers[0]
	}

	p, q := numbers[0], max(numbers[0], numbers[1])
	for i := 2; i < n; i++ {
		p, q = q, max(p+numbers[i], q)
	}
	return q
}


func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}


func main() {
	numbers := []int{2,7,9,3,1}
	fmt.Println(rob(numbers))
}