// LeetCode 主站 Problem Nr. 1305: 两棵二叉搜索树中的所有元素

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	return method1(root1, root2)
}

func method1(root1, root2 *TreeNode) []int {
	var traversal func(n *TreeNode) []int
	traversal = func(n *TreeNode) []int {
		result := make([]int, 0)
		
		if n == nil {
			return result
		}

		result = append(result, traversal(n.Left)...)
		result = append(result, n.Val)
		result = append(result, traversal(n.Right)...)
		return result
	}

	array1, array2 := traversal(root1), traversal(root2)
	length1, length2 := len(array1), len(array2)
	pointer1, pointer2 := 0, 0
	result := make([]int, 0)
	for pointer1 < length1 && pointer2 < length2 {
		if array1[pointer1] < array2[pointer2] {
			result = append(result, array1[pointer1])
			pointer1++
		} else if array1[pointer1] > array2[pointer2] {
			result = append(result, array2[pointer2])
			pointer2++
		} else {
			result = append(result, []int{array1[pointer1], array2[pointer2]}...)
			pointer1++
			pointer2++
		}
	}

	if pointer1 < length1 {
		result = append(result, array1[pointer1:]...)
	}

	if pointer2 < length2 {
		result = append(result, array2[pointer2:]...)
	}

	return result
}