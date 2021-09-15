// LeetCode 主站 Problem Nr. 1008: 前序遍历构造二叉搜索树

package main

import (
	"sort"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type TreeNode = ds.BinaryNode

func bstFromPreorder(preOrder []int) *TreeNode {
	return method1(preOrder)
}

func method1(preOrder []int) *TreeNode {
	inOrder := make([]int, len(preOrder))
	copy(inOrder, preOrder)
	sort.Ints(inOrder)

	mapping := make(map[int]int)
	for i, v := range inOrder {
		mapping[v] = i
	}

	var buildTree func(start, end int) *TreeNode
	buildTree = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}

		rootValue := preOrder[0]
		rootIndex := mapping[rootValue]
		preOrder = preOrder[1:]
		tree := &TreeNode{
			Val: rootValue,
			Left: buildTree(start, rootIndex-1),
			Right: buildTree(rootIndex+1, end),
		}
		return tree
	}

	return buildTree(0, len(preOrder)-1)
}

func main() {
	bstFromPreorder([]int{8,5,1,7,10,12})
}