// 剑指Offer Problem Nr. 09: 用栈实现队列

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type SinglyLinkedNode = ds.SinglyLinkedNode

type Stack struct {
	size int
	sentinel *SinglyLinkedNode
}

func NewStack() *Stack {
	return &Stack{
		size: 0,
		sentinel: &SinglyLinkedNode{
			Val: -1,
		},
	}
}

type CQueue struct {
	input, output *Stack
}

func Constructor() CQueue {
	return CQueue{
		input: NewStack(),
		output: NewStack(),
	}
}

func (cq *CQueue) AppendTail(x int) {
	node := &SinglyLinkedNode{
		Val: x,
		Next: cq.input.sentinel.Next,
	}
	cq.input.sentinel.Next = node
	cq.input.size += 1
}

func (cq *CQueue) DeleteHead() int {
	if cq.output.size == 0 {
		if cq.input.size == 1 {
			return -1
		}
		cq.fromInputToOutput()
	}

	node := cq.output.sentinel.Next
	cq.output.sentinel.Next = cq.output.sentinel.Next.Next
	cq.output.size -= 1
	return node.Val
}

func (cq *CQueue) fromInputToOutput() {
	for cq.input.size > 0 {
		node := cq.input.sentinel.Next
		cq.input.sentinel.Next = cq.input.sentinel.Next.Next
		cq.input.size -= 1

		node.Next = cq.output.sentinel.Next
		cq.output.sentinel.Next = node
		cq.output.size += 1
	}
}

