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
	// ...
	return nil
}