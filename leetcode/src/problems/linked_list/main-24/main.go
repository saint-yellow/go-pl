// LeetCode 主站 Problem Nr. 24: 两两交换链表中的节点

package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)


type ListNode = ds.SinglyLinkedNode

func swapPairs(head *ListNode) *ListNode {
	return method1(head)
}

func method1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	sentinel := &ListNode{Val: -1, Next: head}
	pointer := sentinel
	for pointer.Next != nil && pointer.Next.Next != nil {
		temp := pointer.Next

		pointer.Next = temp.Next
		temp.Next = pointer.Next.Next
		pointer.Next.Next = temp

		pointer = pointer.Next.Next
	}

	return sentinel.Next
}

func method2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	temp := head.Next
	head.Next = method2(temp.Next)
	temp.Next = head
	return temp
}

func main() {
	list := ds.BuildSinglyLinkedList([]int{1,2,3,4,5,6})
	fmt.Println(list.ToList())

	list = swapPairs(list)
	fmt.Println(list.ToList())

	list = method2(list)
	fmt.Println(list.ToList())
}