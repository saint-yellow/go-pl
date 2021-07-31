// LeetCode Problem Nr. 641: 设计循环双端队列

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type DoublyLinkedNode = ds.DoublyLinkedNode

type MyCircularDeque struct {
	capacity, size int
	head, tail *DoublyLinkedNode
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	head := &DoublyLinkedNode{Val: -1}
	tail := &DoublyLinkedNode{Val: -1}
	head.Next = tail
	tail.Previous = head

	return MyCircularDeque{
		capacity: k,
		size: 0,
		head: head,
		tail: tail,
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (cdq *MyCircularDeque) InsertFront(value int) bool {
	if cdq.IsFull() {
		return false
	}

	node := &DoublyLinkedNode{
		Val: value,
		Previous: cdq.head,
		Next: cdq.head.Next,
	}
	cdq.head.Next.Previous = node
	cdq.head.Next = node
	cdq.size += 1
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (cdq *MyCircularDeque) InsertLast(value int) bool {
	if cdq.IsFull() {
		return false
	}

	node := &DoublyLinkedNode{
		Val: value,
		Previous: cdq.tail.Previous,
		Next: cdq.tail,
	}
	cdq.tail.Previous.Next = node
	cdq.tail.Previous = node
	cdq.size += 1
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (cdq *MyCircularDeque) DeleteFront() bool {
	if cdq.IsEmpty() {
		return false
	}

	cdq.head.Next = cdq.head.Next.Next
	cdq.head.Next.Previous = cdq.head
	cdq.size -= 1
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (cdq *MyCircularDeque) DeleteLast() bool {
	if cdq.IsEmpty() {
		return false
	}

	cdq.tail.Previous = cdq.tail.Previous.Previous
	cdq.tail.Previous.Next = cdq.tail
	cdq.size -= 1
	return true
}


/** Get the front item from the deque. */
func (cdq *MyCircularDeque) GetFront() int {
	return cdq.head.Next.Val
}


/** Get the last item from the deque. */
func (cdq *MyCircularDeque) GetRear() int {
	return cdq.tail.Previous.Val
}


/** Checks whether the circular deque is empty or not. */
func (cdq *MyCircularDeque) IsEmpty() bool {
	return cdq.size == 0
}


/** Checks whether the circular deque is full or not. */
func (cdq *MyCircularDeque) IsFull() bool {
	return cdq.size == cdq.capacity
}
