// LeetCode Problem Nr. 226: 反转二叉树

package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	return method1(root)
}

// 基于递归的解法
func method1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	method1(root.Left)
	method1(root.Right)

	return root
}

func method2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	method1(root.Left)
	method1(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

// 基于前序遍历的解法
func method3(root *TreeNode) *TreeNode {
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
			stack = stack[:len(stack)-1]
			node.Left, node.Right = node.Right, node.Left
		}
	}
	return root
}

// 基于后序遍历的解法
func method4(root *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0)
	
	if root != nil {
		stack = append(stack, root)
	}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node != nil {
			stack = stack[:len(stack)-1]

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}

			stack = append(stack, node)
			stack = append(stack, nil)
		} else {
			stack = stack[:len(stack)-1]
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node.Left, node.Right = node.Right, node.Left
		}
	}
	return root
}

// 基于层序遍历的解法
func method5(root *TreeNode) *TreeNode {
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}
