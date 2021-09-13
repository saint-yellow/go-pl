// LeetCode 主站 Problem Nr. 95: 恢复二叉搜索树

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func recoverTree(root *TreeNode) {
	method1(root)
}

func method1(root *TreeNode) {
	var traversal func(node *TreeNode) []*TreeNode
	traversal = func(node *TreeNode) []*TreeNode {
		result := []*TreeNode{}
		if node == nil {
			return result
		}

		result = append(result, traversal(node.Left)...)
		result = append(result, node)
		result = append(result, traversal(node.Right)...)
		return result
	}

	nodes := traversal(root)
	length := len(nodes)
	if length < 2 {
		return
	}

	findSwappedLocations := func(nodes []*TreeNode) []int {
		locations := make([]int, 0)
		for i := 1; i < length; i++ {
			if nodes[i-1].Val > nodes[i].Val {
				locations = append(locations, []int{i-1, i}...)
			}
		}
		return locations
	}

	locations := findSwappedLocations(nodes)
	nodes[locations[0]].Val, nodes[locations[len(locations)-1]].Val = nodes[locations[len(locations)-1]].Val, nodes[locations[0]].Val
}

func main() {
	tree := ds.BuildBinaryTree([]int{1,2,3})
	recoverTree(tree)
}