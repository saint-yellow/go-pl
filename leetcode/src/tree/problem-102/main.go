package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type TreeNode = ds.BinaryNode

func levelOrder(root *TreeNode) [][]int {
	return method1(root)
}

func method1(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}

	result := make([][]int, 0)
	for len(queue) > 0 {
		size := len(queue)
		currentLevel := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, currentLevel)
	}

	return result
}

func method2(root *TreeNode) [][]int {
	result := make([][]int, 0)
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(result) < level+1 {
			result = append(result, make([]int, 0))
		}
		result[level] = append(result[level], node.Val)
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	if root == nil {
		return result
	}

	dfs(root, 0)
	return result
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: nil,
			Right: nil,
		},
	}
	fmt.Println(method1(tree))
	fmt.Println(method2(tree))
}