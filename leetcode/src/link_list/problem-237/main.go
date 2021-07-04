package main

type ListNode struct {
	Value int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Value = node.Next.Value
	node.Next = node.Next.Next
}