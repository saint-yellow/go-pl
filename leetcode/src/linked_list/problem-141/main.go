package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type ListNode = ds.SinglyLinkedNode

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