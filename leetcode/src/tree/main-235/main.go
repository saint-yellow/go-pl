// LeetCode Problem Nr. 235: 二叉搜索树的最近公共祖先

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return method2(root, p, q)
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

func method2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val > p.Val && root.Val > q.Val {
		left := method1(root.Left, p, q)
		if left != nil {
			return left
		}
	}

	if root.Val < p.Val && root.Val < q.Val {
		right := method1(root.Right, p, q)
		if right != nil {
			return right
		}
	}

	return root
}