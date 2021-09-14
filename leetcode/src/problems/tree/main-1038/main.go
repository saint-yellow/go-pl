// LeetCode 主站 Problem Nr. 1038: 把二叉搜索树转换为累加树

package main

import (
	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type TreeNode = ds.BinaryNode

func bstToGst(root *TreeNode) *TreeNode {
	return method1(root)
}

func method1(root *TreeNode) *TreeNode {
	var backtracing func(node *TreeNode, sum *int) *TreeNode
	backtracing = func(node *TreeNode, sum *int) *TreeNode {
		if node == nil {
			return node
		}

		backtracing(node.Right, sum)
		temp := *sum
		*sum += node.Val
		node.Val += temp
		backtracing(node.Left, sum)
		return node
	}

	sum := 0
	return backtracing(root, &sum)
}

func main() {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	bstToGst(tree)
}