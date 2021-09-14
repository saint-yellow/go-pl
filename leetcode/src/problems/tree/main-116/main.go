// LeetCode 主站 Problem Nr. 116: 填充每个节点的下一个右侧节点指针

package main

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	return method1(root)
}

func method1(root *Node) *Node {
	queue := make([]*Node, 0)
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			if i == size-1 {
				queue[i].Next = nil
			} else {
				queue[i].Next = queue[i+1]
			}
		}
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}

func method2(root *Node) *Node {
	if root == nil {
		return root
	}

	for n1 := root; n1.Left != nil; n1 = n1.Left {
		for n2 := n1; n2 != nil; n2 = n2.Next {
			n2.Left.Next = n2.Right

			if n2.Next != nil {
				n2.Right.Next = n2.Next.Left
			}
		}
	}

	return root
}