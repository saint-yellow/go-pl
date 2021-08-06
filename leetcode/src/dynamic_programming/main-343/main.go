// LeetCode 主站 Problem Nr. 343: 整数拆分

package main

func integerBreak(n int) int {
	return method1(n)
}

func method1(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = maxInt(dp[i], maxInt((i-j)*j, dp[i-j]*j))
		}
	}
	return dp[n]
}

func method2(n int) int {
	if n == 2 {
		return 1
	}
	if n ==3 {
		return 2
	}
	if n == 4 {
		return 4
	}

	result := 1
	for n > 4 {
		result *= 3
		n -= 3
	}
	result *= n
	return result
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}