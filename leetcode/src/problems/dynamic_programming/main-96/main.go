// LeetCode 主站 Problem Nr. 96: 不同的二叉搜索树

package main

func numTrees(n int) int {
	return method1(n)
}

func method1(n int) int {
	if n == 0 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 1
	
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}