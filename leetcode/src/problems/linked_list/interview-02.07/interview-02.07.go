// 面试 Problem Nr. 02.07: 链表相交

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type ListNode = ds.SinglyLinkedNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	return method1(headA, headB)
}

func method1(headA, headB *ListNode) *ListNode {
	mapping := make(map[*ListNode]int)
	for headA != nil {
		mapping[headA] = headA.Val
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := mapping[headB]; ok {
			return headB
		}

		headB = headB.Next
	}

	return nil
}

func method2(headA, headB *ListNode) *ListNode {
	pointerA, pointerB := headA, headB
	lenA, lenB := 0, 0

	for pointerA != nil {
		pointerA = pointerA.Next
		lenA++
	}

	for pointerB != nil {
		pointerB = pointerB.Next
		lenB++
	}

	step := 0
	var fast, slow *ListNode

	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}

	for i := 0; i < step; i++ {
		fast = fast.Next
	}

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}