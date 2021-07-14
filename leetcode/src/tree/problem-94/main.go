// LeetCode Problem Nr. 94: 二叉树的中序遍历

package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/tree/ds"
)

type TreeNode = ds.BinaryNode

func inorderTraversal(root *TreeNode) []int {
	return method1(root)
}

func method1(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}

func method2(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}

	if root != nil {
		stack = append(stack, root)
	}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node != nil {
			stack = stack[:len(stack)-1]

			if node.Right != nil {
				stack = append(stack, node.Right)
			}

			stack = append(stack, node)
			stack = append(stack, nil)

			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			stack = stack[:len(stack)-1]
			node = stack[len(stack)-1]
			result = append(result, node.Val)
			stack = stack[:len(stack)-1]
		}
	}

	return result
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: nil,
			Right: nil,
		},
	}
	fmt.Println(inorderTraversal(tree))
}