// LeetCode 主站 Problem Nr. 203: 移除链表元素

package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type ListNode = ds.SinglyLinkedNode

func removeElements(head *ListNode, val int) *ListNode {
	return method1(head, val)
}

// 基于迭代的解法
func method1(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	sentinel := &ListNode{Val: -1, Next: head}
	for pointer := sentinel; pointer.Next != nil; {
		if pointer.Next.Val == val {
			pointer.Next = pointer.Next.Next
		} else {
			pointer = pointer.Next
		}
	}
	return sentinel.Next
}

// 基于递归的解法
func method2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	head.Next = method2(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

func main() {
	list1 := ds.BuildSinglyLinkedList([]int{7, 7, 7, 7})
	fmt.Println(list1.ToList())

	list2 := removeElements(list1, 7)
	fmt.Println(list2.ToList())

	list3 := method2(list1, 6)
	fmt.Println(list3.ToList())
}