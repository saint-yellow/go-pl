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
	startNode := root
	for startNode != nil {
		var nextNode, lastNode *Node
		handle := func(currentNode *Node) {
			if currentNode == nil {
				return
			}

			if nextNode == nil {
				nextNode = currentNode
			}

			if lastNode != nil {
				lastNode.Next = currentNode
			}
			lastNode = currentNode
		}

		for p := startNode; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		startNode = nextNode
	}
	return root
}