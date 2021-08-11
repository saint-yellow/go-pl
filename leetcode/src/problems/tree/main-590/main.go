// LeetCode 主站 Problem Nr. 590: N叉树的后序遍历

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type Node = ds.NAryNode

func postorder(root *Node) []int {
    return method1(root)
}

func method1(root *Node) []int {
	var dfs func(root *Node, result *[]int)
	dfs = func(root *Node, result *[]int) {
		if root == nil {
			return
		}

		if root.Children != nil {
			for _, child := range root.Children {
				dfs(child, result)
			}
		}
		*result = append(*result, root.Val)
	}

	result := make([]int, 0)
	dfs(root, &result)
	return result
}