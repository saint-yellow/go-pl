// LeetCode 主站 Problem Nr. 95: 不同的二叉搜索树 II

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func generateTrees(n int) []*TreeNode {
    return method1(n)
}

func method1(n int) []*TreeNode {
	var backtracking func(start, end int) []*TreeNode
	backtracking = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}

		var result []*TreeNode
		for i := start; i <= end; i++ {
			leftTrees := backtracking(start, i-1)
			rightTrees := backtracking(i+1, end)

			for _, lt := range leftTrees {
				for _, rt := range rightTrees {
					tree := &TreeNode{
						Val: i,
						Left: lt,
						Right: rt,
					}
					result = append(result, tree)
				}
			}
		}
		return result
	}

	return backtracking(1, n)
}

func main() {
	generateTrees(4)
}