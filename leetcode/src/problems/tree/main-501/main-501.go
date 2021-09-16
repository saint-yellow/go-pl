// LeetCode 主站 Problem Nr. 501: 二叉搜索树中的众数

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func findMode(root *TreeNode) []int {
	return method1(root)
}

func method1(root *TreeNode) []int {
	result := make([]int, 0)
	count, max := 1, 1

	var prev *TreeNode
	var traversal func(n *TreeNode)
	traversal = func(n *TreeNode) {
		if n == nil {
			return
		}

		traversal(n.Left)

		if prev != nil && prev.Val == n.Val {
			count += 1
		} else {
			count = 1
		}

		if count >= max {
			if count > max && len(result) > 0 {
				result = []int{n.Val}
			} else {
				result = append(result, n.Val)
			}

			max = count
		}

		prev = n

		traversal(n.Right)
	}

	traversal(root)
	return result
}

func main() {
	findMode(nil)
}