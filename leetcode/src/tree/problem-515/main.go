package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/tree/ds"
)

type TreeNode = ds.BinaryNode

func largestValues(root *TreeNode) []int {
	return method1(root)
}

func method1(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}

	result := make([]int, 0)
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, max(level))
	}

	return result
}

func method2(root *TreeNode) []int {
	result := make([]int, 0)
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(result) < level+1 {
			result = append(result, node.Val)
		}
		
		if result[level] < node.Val {
			result[level] = node.Val
		}

		dfs(node.Left,level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)
	return result
}

func max(digits []int) int {
	max := digits[0]
	for i := 1; i < len(digits); i++ {
		if digits[i] > max {
			max = digits[i]
		}
	}
	return max
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 3,
				Left: nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: nil,
			Right: &TreeNode{
				Val: 9,
				Left: nil,
				Right: nil,
			},
		},
	}

	fmt.Println(method1(tree))
	fmt.Println(method2(tree))
}