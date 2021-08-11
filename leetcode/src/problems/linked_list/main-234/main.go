// LeetCode 主站 Problem Nr. 234: 回文链表

package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type ListNode = ds.SinglyLinkedNode

func isPalindrome(head *ListNode) bool {
	return method2(head)
}

func method1(head * ListNode) bool {
	if head == nil {
		return false
	}

	list := list(head)
	left, right := 0, len(list) - 1
	for left <= right {
		if list[left] != list[right] {
			return false
		}

		left += 1
		right -= 1
	}
	return true
}

func list(head *ListNode) []int {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func method2(head *ListNode) bool {
	// 0
	if head == nil {
		return true
	}

	// 1
	firstHalfEnd := endOfFirstHalf(head)

	// 2
	secondHalfStart := reverse(firstHalfEnd.Next)

	// 3
	pointer1, pointer2 := head, secondHalfStart
	result := true
	for result && pointer2 != nil {
		if pointer1.Val != pointer2.Val {
			result = false
		}

		pointer1 = pointer1.Next
		pointer2 = pointer2.Next
	}

	// 4
	firstHalfEnd.Next = reverse(secondHalfStart)

	// 5
	return result
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var sentinel *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = sentinel
		sentinel = head
		head = next
	}
	return sentinel
}

func main() {
	head := ds.BuildSinglyLinkedList([]int{1, 2, 2, 1})
	fmt.Println(isPalindrome(head))
}