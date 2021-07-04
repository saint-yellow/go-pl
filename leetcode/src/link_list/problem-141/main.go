package main

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
    nodeMap := make(map[*ListNode]int)
    for head != nil {
        if _, ok := nodeMap[head]; ok {
            return true
        }

        nodeMap[head] = 1
        head = head.Next
    }

    return false
}