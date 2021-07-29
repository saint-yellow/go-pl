// LeetCode 主站 Problem Nr. 328: 奇偶链表

package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type ListNode = ds.SinglyLinkedNode

func oddEvenList(head *ListNode) *ListNode {
	return method2(head)
}

func method1(head *ListNode) *ListNode {
	evenSentinel, oddSentinel := &ListNode{Val: -1}, &ListNode{Val: -1}
	evenPointer, oddPointer := evenSentinel, oddSentinel

	count := 0
	for head != nil {
		if count % 2 == 0 {
			evenPointer.Next = &ListNode{Val: head.Val}
			evenPointer = evenPointer.Next
		} else {
			oddPointer.Next = &ListNode{Val: head.Val}
			oddPointer = oddPointer.Next
		}
		count += 1
		head = head.Next
	}

	evenPointer.Next = oddSentinel.Next
	return evenSentinel.Next
}

func method2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	oddHead, evenHead := head, head.Next
	oddPointer, evenPointer := oddHead, evenHead
	for evenPointer != nil && evenPointer.Next != nil {
		oddPointer.Next = evenPointer.Next
		oddPointer = oddPointer.Next
		
		evenPointer.Next = oddPointer.Next
		evenPointer = evenPointer.Next
	}
	oddPointer.Next = evenHead
	return oddHead
}

func main() {
	var list1 *ListNode = nil
	fmt.Println(list1.ToList())
	fmt.Println(oddEvenList(list1).ToList())

	list2 := ds.BuildSinglyLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Println(list2.ToList())
	fmt.Println(oddEvenList(list2).ToList())

	list3 := ds.BuildSinglyLinkedList([]int{2, 1, 3, 5, 6, 4, 7})
	fmt.Println(list3.ToList())
	fmt.Println(oddEvenList(list3).ToList())
}