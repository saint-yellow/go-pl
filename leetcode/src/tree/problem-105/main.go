// LeetCode Problem Nr. 105: 从前序与中序遍历序列构造二叉树

package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return method1(preorder, inorder)
}

func method1(preorder, inorder []int) *TreeNode {
	mapping := make(map[int]int)
	for i, v := range inorder {
		mapping[v] = i
	}

	var builder func(leftIndex, rightIndex int) *TreeNode
	builder = func(leftIndex, rightIndex int) *TreeNode {
		if leftIndex > rightIndex {
			return nil
		}

		value := preorder[0]
		index := mapping[value]
		preorder = preorder[1:]

		root := &TreeNode{
			Val: value,
		}
		root.Left = builder(leftIndex, index-1)
		root.Right = builder(index+1, rightIndex)

		return root

	}

	return builder(0, len(inorder)-1)
}
