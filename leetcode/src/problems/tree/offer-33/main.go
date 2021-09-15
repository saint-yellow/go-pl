// 剑指 Offer Problem Nr. 33: 二叉搜索树的后序遍历序列

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func verifyPostorder(postOrder []int) bool {
	return method2(postOrder, 0, len(postOrder)-1)
}

func method1(postOrder []int) (result bool) {
	return
}

func method2(postOrder []int, start, end int) bool {
	if start > end {
		return true
	}

	p := start
	for postOrder[p] < postOrder[end] {
		p++
	}
	m := p
	for postOrder[p] > postOrder[end] {
		p++
	}
	return p == end && method2(postOrder, start, m-1) && method2(postOrder, m, end-1)
}