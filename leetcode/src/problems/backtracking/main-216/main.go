// LeetCode 主站 Problem Nr. 216: 组合总和 III

package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	return method1(k, n)
}

func method1(k, n int) [][]int {
	result := make([][]int, 0)

	if n <= 0 || k <= 0 || n < k {
		return result
	}

	track := make([]int, 0)

	var backtracking func(n, k , sum, startIndex int)
	backtracking = func(n, k, sum, startIndex int) {
		if sum > n {
			return
		}

		if len(track) == k {
			if sum == n {
				temp := make([]int, k)
				copy(temp, track)
				result = append(result, temp)
			}
			return
		}

		for i := startIndex; i <= 9 - (k - len(track)) + 1; i++ {
			track = append(track, i)
			sum += i
			backtracking(n, k, sum, i + 1)
			track = track[:len(track) - 1]
			sum -= i
		}
	}

	backtracking(n, k, 0, 1)
	return result
}

func main() {
	c := combinationSum3(2, 18)
	fmt.Println(c)
}