// 剑指Offer Problem Nr. 22: 链表中倒数第k个节点

package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type ListNode = ds.SinglyLinkedNode

func getKthFromEnd(head *ListNode, k int) *ListNode {
	return method1(head, k)
}

func method1(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func main() {
	linkedList := ds.BuildSinglyLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Println(getKthFromEnd(linkedList, 2).ToList())
}