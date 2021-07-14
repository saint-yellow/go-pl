package main

import "github.com/saint-yellow/go-pl/leetcode/src/linked_list/ds"

type ListNode = ds.SinglyLinkedNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return method3(head, n)
}

func method1(head *ListNode, n int) *ListNode {
	getLength := func (node *ListNode) int {
		count := 0
		for node != nil {
			count += 1
			node = node.Next
		}
		return count
	}

	length := getLength(head)
	sentinel := &ListNode{
		Val: -1,
		Next: head,
	}
	current := sentinel
	for i := 1; i < length+1-n; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	return sentinel.Next
}

func method2(head *ListNode, n int) *ListNode {
	sentinel := &ListNode{
		Val: -1,
		Next: head,
	}
	stack := make([]*ListNode, 0)
	current := sentinel
	for current!= nil {
		stack = append(stack, current)
		current = current.Next
	}
	current = stack[len(stack)-1-n]
	current.Next = current.Next.Next
	return sentinel.Next
}

func method3(head *ListNode, n int) *ListNode {
	sentinel := &ListNode{
		Val: -1,
		Next: head,
	}
	slow, fast := sentinel, sentinel.Next
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return sentinel.Next
}