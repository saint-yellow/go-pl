// LeetCode 主站 Problem Nr. 40: 组总和 II

package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	return method1(candidates, target)
}

func method1(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	track := make([]int, 0)
	used := make([]bool, len(candidates))

	var backtracking func(n, target, sum, startIndex int)
	backtracking = func(n, target, sum, startIndex int) {
		if sum > target {
			return
		}

		if sum == target {
			temp := make([]int, len(track))
			copy(temp, track)
			result = append(result, temp)
			return
		}

		for i := startIndex; i < n; i++ {
			if i > 0 && candidates[i] == candidates[i - 1] && !used[i - 1] {
				continue
			}

			used[i] = true
			track = append(track, candidates[i])
			sum += candidates[i]

			backtracking(n, target, sum, i + 1)

			used[i] = false
			track = track[:len(track)-1]
			sum -= candidates[i]
		}
	}

	n := len(candidates)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	backtracking(n, target, 0, 0)
	return result
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	c := combinationSum2(candidates, target)
	fmt.Println(c)
}