// LeetCode 主站 Problem Nr. 2: 两数相加

package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)


type ListNode = ds.SinglyLinkedNode

func method1(l1, l2 *ListNode) *ListNode {
	c := 0
	head := new(ListNode)
	l3 := head
	for l1 != nil || l2 != nil || c != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		v3 := v1+v2+c
		if v3 >= 10 {
			c = 1
			v3 -= 10
		} else {
			c = 0
		}
		fmt.Println(v3)
		l3.Next = &ListNode{Val: v3, Next: nil}
		l3 = l3.Next
	}
	return head.Next
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	fmt.Println(method1(l1, l2).ToList())
}