// LeetCode 主站 Problem Nr. 77: 组合

package main

import "fmt"

func combine(n, k int) [][]int {
	return method1(n, k)
}

// 基于回溯的解法
func method1(n, k int) [][]int {
	result := make([][]int, 0) 
	path := make([]int, 0)
	if n <= 0 || k <= 0 || k > n {
		return result
	}

	var backtracking func(n, k, startIndex int)
	backtracking = func(n, k, startIndex int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := startIndex; i <= n - (k - len(path)) + 1; i++ {
			path = append(path, i)
			backtracking(n, k, i + 1)
			path = path[0:len(path) - 1]
		}
	}

	backtracking(n, k, 1)
	return result
}

func main() {
	c := combine(4, 3)
	fmt.Println(c)
}