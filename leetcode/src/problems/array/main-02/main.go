package main

import "fmt"

type ListNode struct {
	Value int
	Next *ListNode
}

func newListNode(value int, next *ListNode) *ListNode {
	return &ListNode{Value: value, Next: next}
}

func (n *ListNode) listAllValues() []int {
	if n == nil {
		return make([]int, 0)
	}
	return append([]int{n.Value}, n.Next.listAllValues()...)
}

func method1(l1, l2 *ListNode) *ListNode {
	c := 0
	head := new(ListNode)
	l3 := head
	for l1 != nil || l2 != nil || c != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Value
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Value
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
		l3.Next = &ListNode{Value: v3, Next: nil}
		l3 = l3.Next
	}
	return head.Next
}

func main() {
	l1 := newListNode(2, newListNode(4, newListNode(3, nil)))
	l2 := newListNode(5, newListNode(6, newListNode(4, nil)))
	fmt.Println(method1(l1, l2).listAllValues())
}