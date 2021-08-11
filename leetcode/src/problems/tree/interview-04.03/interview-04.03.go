package main

import (
	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type ListNode = ds.SinglyLinkedNode
type TreeNode = ds.BinaryNode

func listOfDepth(tree *TreeNode) []*ListNode {
	return method1(tree)
}

func method1(tree *TreeNode) []*ListNode {
	if tree == nil {
		return make([]*ListNode, 0)
	}

	queue := []*TreeNode{tree}
	result := make([]*ListNode, 0)

	for len(queue) > 0 {
		size := len(queue)
		sentinel := &ListNode{Val: -1}
		pointer := sentinel
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			pointer.Next = &ListNode{Val: node.Val}
			pointer = pointer.Next

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, sentinel.Next)
	}

	return result
}


func main() {
}