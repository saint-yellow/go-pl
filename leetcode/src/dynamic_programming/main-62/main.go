// LeetCode Problem Nr. 62: 不同路径

package main

import "fmt"

func uniquePaths(m int, n int) int {
	return method1(m, n)
}

func method1(m, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	dp := make([][]int, m)
	dp[0] = make([]int, n)
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		row := make([]int, n)
		row[0] = 1
		dp[i] = row
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func method2(m, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(method2(3, 7))
}