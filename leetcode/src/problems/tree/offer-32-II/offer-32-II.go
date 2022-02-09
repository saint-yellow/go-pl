package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type TreeNode = ds.BinaryNode

func levelOrder(root *TreeNode) [][]int {
	return method1(root)
}

// 基于广度优先遍历的解法
func method1(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		currentLevel := make([]int, len(queue))
		for i := 0; i < len(queue); i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevel[i] = node.Val
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

// 基于深度优先遍历的解法
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
		dfs(node.Left,level+1)
		dfs(node.Right, level+1)
	}

	if root == nil {
		return result
	}
	dfs(root, 0)
	return result
}

func main() {
	array := make([]int, 2)
	array[0] = 1
	fmt.Println(array)
}