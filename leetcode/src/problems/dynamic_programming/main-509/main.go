// LeetCode 主站 Problem Nr. 509: 斐波那契数

package main

import "fmt"

func fib(n int) int {
	return method1(n)
}

func method1(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func method2(n int) int {
	if n <= 1 {
		return n
	}

	return method1(n-1) + method1(n-2)
}

func method3(n int) int {
	if n <= 1 {
		return n
	}

	p, q := 0, 1
	for i := 2; i <= n; i++ {
		p, q = q, p+q
	}
	return q
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(fib(i), method3(i))
	}
}