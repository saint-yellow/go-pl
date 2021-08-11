// LeetCode 主站 Problem Nr. 206: 反转链表

package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type ListNode = ds.SinglyLinkedNode

func reverseList(head *ListNode) *ListNode {
	return method2(head)
}

// 基于迭代的解法
func method1(head *ListNode) *ListNode {
	var sentinel *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = sentinel
		sentinel, head = head, next
	}
	return sentinel
}

// 基于递归的解法
func method2(head * ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reversedHead := method2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reversedHead
}

func method3(previous, current *ListNode) *ListNode {
	if current == nil {
		return current
	}

	next := current.Next
	current.Next = previous
	return method3(current, next)
}

func main() {
	values := []int{1, 2, 3, 4, 5}
	list := ds.BuildSinglyLinkedList(values)
	fmt.Println(list.ToList())
	fmt.Println(reverseList(list).ToList())
}