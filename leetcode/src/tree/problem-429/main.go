package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/tree/ds"
)

type Node = ds.NAryNode

func levelOrder(root *Node) [][]int {
	return method1(root)
}

func method1(root *Node) [][]int {
	queue := make([]*Node, 0)

	if root != nil {
		queue = append(queue, root)
	}

	result := make([][]int, 0)
	
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Children != nil {
				queue = append(queue, node.Children...)
			}
		}
		result =append(result, level)
	}

	return result
}

func method2(root *Node) [][]int {
	result := make([][]int, 0)
	var dfs func(node *Node, level int)
	dfs = func(node *Node, level int) {
		if node == nil {
			return
		}

		if len(result) < level+1 {
			result = append(result, make([]int, 0))
		}
		result[level] = append(result[level], node.Val)
		if node.Children != nil {
			for _, c := range node.Children {
				dfs(c, level+1)
			}
		}
	}

	dfs(root, 0)
	return result
}

func main() {
	tree := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
						Children: []*Node{},
					},
					{
						Val: 6,
						Children: []*Node{},
					},
				},
			},
			{
				Val: 2,
				Children: nil,
			},
			{
				Val: 4,
				Children: nil,
			},
		},
	}

	fmt.Println(method1(tree))
	fmt.Println(method2(tree))
}

