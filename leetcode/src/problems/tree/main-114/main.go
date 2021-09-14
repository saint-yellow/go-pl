// LeetCode 主站 Problem Nr. 114: 二叉树展开为链表

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func flatten(root *TreeNode) {
	method1(root)
}

func method1(root *TreeNode) {
	var traversal func(node *TreeNode) []*TreeNode
	traversal = func(node *TreeNode) []*TreeNode {
		result := make([]*TreeNode, 0)
		if node == nil {
			return result
		}

		result = append(result, node)
		result = append(result, traversal(node.Left)...)
		result = append(result, traversal(node.Right)...)
		return result
	}

	preOrder := traversal(root)
	n := len(preOrder)
	if n == 0 {
		return
	}

	for i := 1; i < n; i++ {
		preOrder[i - 1].Left = nil
		preOrder[i - 1].Right = preOrder[i]
	}
	
}

func main() {
	tree := &TreeNode{
		Val: 0,
	}
	flatten(tree)
}