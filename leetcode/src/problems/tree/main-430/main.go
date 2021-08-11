// LeetCode 主站 Problem Nr. 430: 扁平化多级双向链表

package main

type Node struct {
	Val int
	Prev, Next, Child *Node
}

func flatten(root *Node) *Node {
	return method1(root)
}


// 线性级空间复杂度的解法
func method1(root *Node) *Node {
	sentinel := &Node{Val: -1, Next: root}
	slow, fast := sentinel, sentinel.Next

	// 存储转折点
	turningPoints := make([]*Node, 0)

	for fast != nil {
		if fast.Child != nil {
			slow = fast
			fast = fast.Child

			turningPoints = append(turningPoints, slow.Next)

			slow.Next = fast
			fast.Prev = slow
			slow.Child = nil
		} else {
            slow = fast
		    fast = fast.Next
        }
	}

    fast = slow

	n := len(turningPoints)
	for i := n - 1; i >= 0; i-- {
		node := method1(turningPoints[i])
		fast.Next = node
        if node != nil {
            node.Prev = fast
        }

		for fast.Next != nil {
			fast = fast.Next
		}
	}

	return sentinel.Next
}

// 常数级空间复杂度的解法
func method2(root *Node) *Node {
	previous := &Node{Val: -1, Next: root}
	for current := root; current != nil; current = current.Next {
		if current.Child != nil {
			temp := current.Child
			for temp.Next != nil {
				temp = temp.Next
			}
			if current.Next != nil {
				current.Next.Prev = temp
			}
			temp.Next = current.Next
			current.Child.Prev = current
			current.Next = current.Child
			current.Child = nil
		}
	}
	return previous.Next
}