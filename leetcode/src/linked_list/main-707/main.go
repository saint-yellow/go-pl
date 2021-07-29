package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type ListNode = ds.DoublyLinkedNode

type MyLinkedList struct {
	Size int
	head *ListNode
	tail *ListNode
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	head := &ListNode{Val:-1}
	tail := &ListNode{Val:-1}
	head.Next = tail
	tail.Previous = head

	return MyLinkedList{
		Size: 0,
		head: head,
		tail: tail,
	}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (mll *MyLinkedList) Get(index int) int {
	if index < 0 || index >= mll.Size {
		return -1
	}

	pointer := mll.head.Next
	for i := 0; i < index; i++ {
		pointer = pointer.Next
	}

	return pointer.Val
}



//  Add a node of value val before the first element of the linked list. 
//  After the insertion, the new node will be the first node of the linked list. 
func (mll *MyLinkedList) AddAtHead(val int)  {
	node := &ListNode{
		Val: val,
		Previous: mll.head,
		Next: mll.head.Next,
	}

	mll.head.Next = node
	node.Next.Previous = node

	mll.Size += 1
}

// Append a node of value val to the last element of the linked list.
func (mll *MyLinkedList) AddAtTail(val int)  {
	node := &ListNode{
		Val: val,
		Previous: mll.tail.Previous,
		Next: mll.tail,
	}

	mll.tail.Previous = node
	node.Previous.Next = node

	mll.Size += 1
}


// Add a node of value val before the index-th node in the linked list. 
// If index equals to the length of linked list, the node will be appended to the end of linked list. 
// If index is greater than the length, the node will not be inserted.
func (mll *MyLinkedList) AddAtIndex(index int, val int)  {
	if index < 0 || index > mll.Size {
		return
	}

	pointer := mll.head
	for i := 0; i < index; i++ {
		pointer = pointer.Next
	}

	// new node to be added
	node := &ListNode{
		Val: val,
		Previous: pointer,
		Next: pointer.Next,
	}
	pointer.Next.Previous = node
	pointer.Next = node

	mll.Size += 1
}


// Delete the index-th node in the linked list, if the index is valid.
func (mll *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index >= mll.Size {
		return
	}

	pointer := mll.head
	for i := 0; i < index; i++ {
		pointer = pointer.Next
	}

	// old node to be deleted
	// node := pointer.Next
	// pointer.Next = node.Next
	// node.Next.Previous = pointer

	// adjust pointers
	pointer.Next = pointer.Next.Next
	pointer.Next.Previous = pointer

	mll.Size -= 1
}

func (mll *MyLinkedList) ToList(reverse bool) []int {
	if reverse {
		return mll.tail.Previous.ToList(reverse)
	}

	return mll.head.Next.ToList(reverse)
}

func main() {
	linkedList := Constructor()
	fmt.Println(linkedList.ToList(false))
	
	linkedList.AddAtHead(7)
	linkedList.AddAtHead(2)
	linkedList.AddAtHead(1)
	fmt.Println(linkedList.ToList(false))

	linkedList.AddAtIndex(3, 0)
	fmt.Println(linkedList.ToList(false))

	linkedList.DeleteAtIndex(2)
	fmt.Println(linkedList.ToList(false))

	linkedList.AddAtHead(6)
	fmt.Println(linkedList.ToList(false))

	linkedList.AddAtTail(4)
	fmt.Println(linkedList.ToList(false))

	fmt.Println(linkedList.Get(4))
	fmt.Println(linkedList.ToList(false))

	linkedList.AddAtHead(4)
	fmt.Println(linkedList.ToList(false))

	linkedList.AddAtIndex(5, 0)
	fmt.Println(linkedList.ToList(false))

	linkedList.AddAtHead(6)
	fmt.Println(linkedList.ToList(false))
}