package main

import "fmt"

type ListNode struct {
	Value int
	Next *ListNode
}

func (node *ListNode) toList() []int {
	result := make([]int, 0)
	for node != nil {
		result = append(result, node.Value)
		node = node.Next
	}
	return result
}

func deleteNode(node *ListNode) {
	method1(node)
}

func method1(node *ListNode) {
	node.Value = node.Next.Value
	node.Next = node.Next.Next
}

func main() {
	list := &ListNode{
		Value: 1,
		Next: &ListNode{
			Value: 2,
			Next: &ListNode{
				Value: 3,
				Next: nil,
			},
		},
	}
	fmt.Println(list.toList())
	deleteNode(list)
	fmt.Println(list.toList())
}