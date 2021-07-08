package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	return method2(root)
}

func method1(root *TreeNode) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		currentSize := len(queue)
		currentLevel := make([]int, 0)
		for i := 0; i < currentSize; i++ {
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

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
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

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

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