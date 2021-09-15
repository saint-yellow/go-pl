// LeetCode 主站 Problem Nr. 897: 递增顺序搜索树

package main

import (
	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type TreeNode = ds.BinaryNode

func increasingBST(root *TreeNode) *TreeNode {
	return method1(root)
}

func method1(root *TreeNode) *TreeNode {
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

	numbers := traversal(root)
	sentinel := &TreeNode{Val: -1}
	pointer := sentinel
	for _, n := range numbers {
		pointer.Right = &TreeNode{Val: n}
		pointer = pointer.Right
	}
	return sentinel.Right
}

func main() {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	increasingBST(tree)
}