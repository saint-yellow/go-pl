// LeetCode 主站 Problem Nr. 70: 爬楼梯

package main

import "fmt"

func climbStairs(n int) int {
	return method1(n)
}

func method1(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func method2(n int) int {
	if n <= 1 {
		return n
	}

	return method2(n-1) + method2(n-2)
}

func method3(n int) int {
	if n <= 1 {
		return n
	}

	p, q := 1, 1
	for i := 2; i <= n; i++ {
		p, q = q, p+q
	}
	return q
}

func main() {
	fmt.Println(climbStairs(7))
}