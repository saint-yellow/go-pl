// LeetCode 主站 Problem Nr. 2: 两数相加

package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type ListNode = ds.SinglyLinkedNode

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	return method1(l1, l2)
}

func method1(l1, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{Val: 0}
	l3 := head
	for l1 != nil || l2 != nil || carry != 0 {
		v1, v2, v3 := 0, 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		s := v1 + v2 + carry
		if s >= 10 {
			carry = 1
			v3 = s - 10
		} else {
			carry = 0
			v3 = s
		}
		l3.Next = &ListNode{Val: v3}
		l3 = l3.Next
	}
	return head.Next
}

func main() {
	l1 := ds.BuildSinglyLinkedList([]int{2, 4, 3}) // 342
	l2 := ds.BuildSinglyLinkedList([]int{7, 0, 8}) // 807

	l3 := addTwoNumbers(l1, l2) // 342 + 807 = 1149

	fmt.Println(l1.ToList())
	fmt.Println(l2.ToList())
	fmt.Println(l3.ToList())
}