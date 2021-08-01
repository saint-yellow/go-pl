// LeetCode 主站 Problem Nr. 39: 组合总和

package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	return method1(candidates, target)
}

func method1(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	track := make([]int, 0)

	var backtracking func(candidates []int, target, sum, startIndex int)		
	backtracking = func(candidates []int, target, sum, startIndex int) {
		if sum > target {
			return
		}

		if sum == target {
			temp := make([]int, len(track))
			copy(temp, track)
			result = append(result, temp)
			return
		}

		for i := startIndex; i < len(candidates); i++ {
			sum += candidates[i]
			track = append(track, candidates[i])
			backtracking(candidates, target, sum, i)
			sum -= candidates[i]
			track = track[:len(track)-1]
		}
	}

	backtracking(candidates, target, 0, 0)
	return result
}

func method2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	track := make([]int, 0)

	var backtracking func(candidates []int, target, sum, startIndex int)		
	backtracking = func(candidates []int, target, sum, startIndex int) {

		if sum == target {
			temp := make([]int, len(track))
			copy(temp, track)
			result = append(result, temp)
			return
		}

		for i := startIndex; i < len(candidates) && sum + candidates[i] <= target; i++ {
			sum += candidates[i]
			track = append(track, candidates[i])
			backtracking(candidates, target, sum, i)
			sum -= candidates[i]
			track = track[:len(track)-1]
		}
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] <= candidates[j]
	})
	backtracking(candidates, target, 0, 0)
	return result
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 4))
	fmt.Println(method2([]int{2, 5, 3}, 4))
}