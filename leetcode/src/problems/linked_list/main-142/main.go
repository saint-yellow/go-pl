package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type ListNode = ds.SinglyLinkedNode

func detectCycle(head *ListNode) *ListNode {
    return method1(head)
}

func method1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil || fast.Next.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = head
	for fast.Next != slow {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func method2(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			p, q := head, slow

			for p != q {
				p = p.Next
				q = q.Next
			}
			
			return p
		}
	}

	return nil
}