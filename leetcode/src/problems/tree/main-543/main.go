// LeetCode 主站 Problem Nr. 543: 二叉树的直径

package main

import (
	"sort"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type TreeNode = ds.BinaryNode

func diameterOfBinaryTree(root *TreeNode) int {
	return method1(root)
}

func method1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var maxDepth func(root *TreeNode) int 
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		depths := make([]int, 2)
		depths[0] = maxDepth(root.Left)
		depths[1] = maxDepth(root.Right)
		sort.Slice(depths, func(i, j int) bool {
			return depths[i] < depths[j]
		})
		return depths[1] + 1
	}

	diameters := make([]int, 3)
	diameters[0] = maxDepth(root.Left) + maxDepth(root.Right)
	diameters[1] = method1(root.Left)
	diameters[2] = method1(root.Right)
	sort.Slice(diameters, func(i, j int) bool {
		return diameters[i] < diameters[j]
	})
	return diameters[2]
}