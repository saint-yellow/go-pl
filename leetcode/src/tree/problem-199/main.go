package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/tree/ds"
)

type TreeNode = ds.BinaryNode

func New(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{val, left, right}
}

func rightSideView(root *TreeNode) []int {
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
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == size-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

	}

	return result
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: nil,
			Right: &TreeNode{
				Val: 5,
				Left: nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: nil,
			Right: &TreeNode{
				Val: 4,
				Left: nil,
				Right: nil,
			},
		},
		
	}
	fmt.Println(method1(tree))
}