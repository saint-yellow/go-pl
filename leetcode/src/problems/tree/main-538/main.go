// LeetCode 主站 Problem Nr. 538: 吧二叉搜索树转换为累加树

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	return method1(root, &sum)
}

func method1(root *TreeNode, sum *int) *TreeNode {
	if root == nil {
		return root
	}

	method1(root.Right, sum)
	temp := *sum
	*sum += root.Val
	root.Val += temp
	method1(root.Left, sum)
	return root
}