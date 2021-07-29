package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type ListNode = ds.SinglyLinkedNode

func hasCycle(head *ListNode) bool {
    return method2(head)
}

func method1(head *ListNode) bool {
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

func method2(head *ListNode) bool {
    if head == nil {
        return false
    }

    slow, fast := head, head.Next
    for slow != fast {
        // 快指针无法移动，即表示链表无环
        if fast == nil || fast.Next == nil || fast.Next.Next == nil {
            return false
        }

        slow = slow.Next
        fast = fast.Next.Next
    }

    // 两指针相会， 即表示链表有环
    return true
}