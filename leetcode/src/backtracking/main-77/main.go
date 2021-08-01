// LeetCode 主站 Problem Nr. 77: 组合

package main

import "fmt"

var (
	result [][]int
	path []int
)

func backtracking(n, k, startIndex int) {
	if len(path) == k {
		temp := make([]int, 2)
		copy(temp, path)
		result = append(result, temp)
		return
	}

	for i := startIndex; i <= n; i++ {
		path = append(path, i)
		backtracking(n, k, i + 1)
		path = path[0:len(path) - 1]
	}
}

func combine(n, k int) [][]int {
	return method1(n, k)
}

func method1(n, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return result
	}
	backtracking(n, k, 1)
	return result
}

func main() {
	combine(1, 1)
	fmt.Println(result)
}