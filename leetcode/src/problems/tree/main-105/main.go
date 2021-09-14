// LeetCode Problem Nr. 105: 从前序与中序遍历序列构造二叉树

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func buildTree(preOrder []int, inOrder []int) *TreeNode {
	return method1(preOrder, inOrder)
}

func method1(preOrder, inOrder []int) *TreeNode {
	mapping := make(map[int]int)
	for i, v := range inOrder {
		mapping[v] = i
	}

	var builder func(leftIndex, rightIndex int) *TreeNode
	builder = func(leftIndex, rightIndex int) *TreeNode {
		if leftIndex > rightIndex {
			return nil
		}

		value := preOrder[0]
		index := mapping[value]
		preOrder = preOrder[1:]

		root := &TreeNode{
			Val: value,
		}
		root.Left = builder(leftIndex, index-1)
		root.Right = builder(index+1, rightIndex)

		return root

	}

	return builder(0, len(inOrder)-1)
}
