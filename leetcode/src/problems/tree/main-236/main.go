// LeetCode Problem Nr. 236: 二叉树的最近公共祖先

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return method1(root, p, q)
}

func method1(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}

	left := method1(root.Left, p, q)
	right := method1(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil && right == nil {
		return left
	}

	if left == nil && right != nil {
		return right
	}

	return nil
}