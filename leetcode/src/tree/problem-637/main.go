package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func New(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{val, left, right}
}

func averageOfLevels(root *TreeNode) []float64 {
	return method1(root)
}

func method1(root *TreeNode) []float64 {
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}

	result := make([]float64, 0)
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

		result = append(result, average(currentLevel))
	}

	return result
}

func average(digits []int) float64 {
	sum := 0
	for _, d := range digits {
		sum += d
	}
	return float64(sum) / float64(len(digits))
}

func main() {
	tree := New(3, New(9, nil, nil), New(20, New(15, nil, nil), New(7, nil, nil)))
	fmt.Println(method1(tree))
}