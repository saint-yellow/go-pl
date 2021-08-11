// LeetCode Problem Nr. 746: 使用最小花费爬楼梯

package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	return method1(cost)
}

func method1(cost []int) int {
	n := len(cost)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return cost[0]
	}

	dp := make([]int, n)
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < n; i++ {
		dp[i] = minInt(dp[i-1], dp[i-2]) + cost[i]
	}
	fmt.Printf("dp   = %v\n", dp)
	return minInt(dp[n-1], dp[n-2])
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func method2(cost []int) int {
	n := len(cost)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return cost[0]
	}

	p, q := cost[0], cost[1]
	for i := 2; i < n; i++ {
		p, q = q, minInt(p, q) + cost[i]
	}
	return minInt(p, q)
}

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Printf("cost = %v\n", cost)
	fmt.Println(minCostClimbingStairs(cost), method2(cost))
}