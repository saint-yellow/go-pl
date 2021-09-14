// LeetCode 主站 Problem Nr. 897: 递增顺序搜索树

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func increasingBST(root *TreeNode) *TreeNode {
	return method1(root)
}