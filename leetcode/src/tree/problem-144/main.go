// LeetCode Problem Nr. 144: 二叉树的前序遍历

package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/tree/ds"
)

type TreeNode = ds.BinaryNode

func preorderTraversal(root *TreeNode) []int {
	return method2(root)
}

func method1(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}

func method2(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

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

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			stack = append(stack, node)
			stack = append(stack, nil)
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
	fmt.Println(preorderTraversal(tree))
}