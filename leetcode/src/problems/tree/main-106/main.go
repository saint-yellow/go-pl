// LeetCode Problem Nr. 106: 从中序与后序遍历序列构造二叉树

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func buildTree(inOrder []int, postOrder []int) *TreeNode {
	return method1(inOrder, postOrder)
}

func method1(inOrder, postOrder []int) *TreeNode {
	mapping := make(map[int]int)
	for i, v := range inOrder {
		mapping[v] = i
	}

	var helperFunc func(leftIndex, rightIndex int) *TreeNode
	helperFunc = func(leftIndex, rightIndex int) *TreeNode {
		if leftIndex > rightIndex {
			return nil
		}

		value := postOrder[len(postOrder)-1]
		index := mapping[value]
		postOrder = postOrder[:len(postOrder)-1]
		root := &TreeNode{
			Val: value,
		}
		root.Right = helperFunc(index+1, rightIndex)
		root.Left = helperFunc(leftIndex, index-1)
		return root
	}
	return helperFunc(0, len(inOrder)-1)
}
