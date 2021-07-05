// LeetCode Problem Nr. 106: 从中序与后序遍历序列构造二叉树

package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	mapping := make(map[int]int)
	for i, v := range inorder {
		mapping[v] = i
	}

	var helperFunc func(leftIndex, rightIndex int) *TreeNode
	helperFunc = func(leftIndex, rightIndex int) *TreeNode {
		if leftIndex > rightIndex {
			return nil
		}

		value := postorder[len(postorder)-1]
		index := mapping[value]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{
			Val: value,
		}
		root.Right = helperFunc(index+1, rightIndex)
		root.Left = helperFunc(leftIndex, index-1)
		return root
	}
	return helperFunc(0, len(inorder)-1)
}
