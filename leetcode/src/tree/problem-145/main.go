// LeetCode Problem Nr. 145: 二叉树的后序遍历

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	return method1(root)
}

func method1(root *TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	result = append(result, postorderTraversal(root.Left)...)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)
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

			stack = append(stack, node)
			stack = append(stack, nil)

			if node.Right != nil {
				stack = append(stack, node.Right)
			}

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
	fmt.Println(method1(tree))
	fmt.Println(method2(tree))
}