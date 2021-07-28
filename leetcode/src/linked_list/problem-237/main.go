package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type ListNode = ds.SinglyLinkedNode

func deleteNode(node *ListNode) {
	method1(node)
}

func method1(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: nil,
			},
		},
	}
	fmt.Println(list.ToList())
	deleteNode(list)
	fmt.Println(list.ToList())
}