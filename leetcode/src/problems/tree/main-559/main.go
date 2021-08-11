// LeetCode 主站 Problem Nr. 559: N叉树的最大深度

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type Node = ds.NAryNode

func maxDepth(root *Node) int {
	return method1(root)
}

func method1(root *Node) int {
	if root == nil {
		return 0
	}

	if root.Children != nil {
		maxChildDepth := 0
		for _, child := range root.Children {
			depth := method1(child)
			if depth > maxChildDepth {
				maxChildDepth = depth
			}
		}
		return maxChildDepth + 1
	}

	return 1
}