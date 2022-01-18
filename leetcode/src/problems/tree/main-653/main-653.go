// LeetCode 主站 Problem Nr. 653: 两数之和 IV - 输入 BST

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func findTarget(root *TreeNode, k int) bool {
	return method1(root, k)
}

func method1(root *TreeNode, k int) bool {
	var traversal func(n *TreeNode) []int
	traversal = func(n *TreeNode) []int {
		result := make([]int, 0)

		if n == nil {
			return result
		}

		result = append(result, traversal(n.Left)...)
		result = append(result, n.Val)
		result = append(result, traversal(n.Right)...)
		return result
	}

	inOrder := traversal(root)
	left, right := 0, len(inOrder)-1
	for left < right {
		sum := inOrder[left] + inOrder[right]
		if sum == k {
			return true
		} else if sum > k {
			right--
		} else {
			left++
		}
	}
	return false
}

func main() {
	findTarget(nil, 0)
}