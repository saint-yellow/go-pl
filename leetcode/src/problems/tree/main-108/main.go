// LeetCode Problem Nr. 108: 将有序数组转换为二叉搜索树

package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type TreeNode = ds.BinaryNode

func sortedArrayToBST(numbers []int) *TreeNode {
	return method1(numbers)
}

func method1(numbers []int) *TreeNode {
	var buildBST func(left, right int) *TreeNode
	buildBST = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		middle := left + (right - left) / 2
		root := &TreeNode{Val: numbers[middle]}
		root.Left = buildBST(left, middle)
		root.Right = buildBST(middle + 1, right)
		return root
	}

	return buildBST(0, len(numbers) - 1)
}

func method2(numbers []int) *TreeNode {
	n := len(numbers)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: -1}
	nodeQueue := []*TreeNode{root}
	leftQueue := []int{0} 
	rightQueue := []int{n - 1}

	for len(nodeQueue) > 0 {
		currentNode := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		left := leftQueue[0]
		leftQueue = leftQueue[1:]

		right := rightQueue[0]
		rightQueue = rightQueue[1:]

		middle := left + (right - left) / 2

		currentNode.Val = numbers[middle]

		if left <= middle - 1 {
			currentNode.Left = &TreeNode{Val: -1}
			nodeQueue = append(nodeQueue, currentNode.Left)
			leftQueue = append(leftQueue, left)
			rightQueue = append(rightQueue, middle - 1)
		}

		if right >= middle + 1 {
			currentNode.Right = &TreeNode{Val: -1}
			nodeQueue = append(nodeQueue, currentNode.Right)
			leftQueue = append(leftQueue, middle + 1)
			rightQueue = append(rightQueue, right)
		}
	}

	return root
}

func main() {
	fmt.Println("hello")
}