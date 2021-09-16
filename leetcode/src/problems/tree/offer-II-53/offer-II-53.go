// 剑指 Offer II Problem Nr. 53: 二叉搜索树中的中序后继

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func inOrderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    return method1(root, p)
}

func method1(root, p *TreeNode) *TreeNode {
	flattenTree := func(n *TreeNode) *TreeNode {
		var traversal func(n *TreeNode) []*TreeNode
		traversal = func(n *TreeNode) []*TreeNode {
			result := make([]*TreeNode, 0)
		
			if n == nil {
				return result
			}

			left, right := n.Left, n.Right
			result = append(result, traversal(left)...)
			n.Left, n.Right = nil, nil
			result = append(result, n)
			result = append(result, traversal(right)...)
			return result
		}

		nodes := traversal(root)
		sentinel := &TreeNode{}
		pointer := sentinel
		for _, n := range nodes {
			pointer.Right = n
			pointer = pointer.Right
		}
		return sentinel.Right
	}

	tree := flattenTree(root)
	for node := tree; node.Right != nil; node = node.Right {
		if node == p {
			return node.Right
		}
	}

	return nil
}

func method2(root, p *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}

	// flag := false
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
			stack = stack[:len(stack)-1]
		}
	}
	return nil
}

func method3(root, p *TreeNode) *TreeNode {
	flag := false
	var traversal func(n, p *TreeNode) *TreeNode
	traversal = func(n, p *TreeNode) *TreeNode {
		if n == nil {
			return n
		}

		if n == p {
			flag = true
		}

		left := traversal(n.Left, p)
		if flag && n.Val > p.Val {
			flag = false
			return n
		}
		right := traversal(n.Right, p)
		if left == nil {
			return right
		}
		return left
	}

	return traversal(root, p)
}