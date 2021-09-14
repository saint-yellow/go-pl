// LeetCode 主站 Problem Nr. 938: 二叉搜索树的范围和

package main

import (
	"fmt"

	"github.com/saint-yellow/go-pl/leetcode/src/ds"
)

type TreeNode = ds.BinaryNode

func rangeSumBST(root *TreeNode, low int, high int) int {
	return method1(root, low, high)
}

func method1(root *TreeNode, low, high int) int {
	var traversal func(node *TreeNode) []int
	traversal = func(node *TreeNode) []int {
		result := make([]int, 0)

		if node == nil {
			return result
		}

		result = append(result, traversal(node.Left)...)
		result = append(result, node.Val)
		result = append(result, traversal(node.Right)...)
		return result
	}

	numbers := traversal(root)
	sum := 0
	for _, value := range numbers {
		if value >= low && value <= high {
			sum += value
		}
	}
	return sum
}

func method2(root *TreeNode, low, high int) (sum int) {
	if root == nil {
		return
	}

	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}
	sum += method2(root.Left, low, high)
	sum += method2(root.Right, low, high)
	return
}

func method3(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}

	if root.Val > high {
		return method3(root.Left, low, high)
	}
	if root.Val < low {
		return method3(root.Right, low, high)
	}
	return root.Val + method3(root.Left, low, high) + method3(root.Right, low, high)
}

func main() {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	fmt.Println(rangeSumBST(tree, 7, 15))
	fmt.Println(method2(tree, 7, 15))
	fmt.Println(method3(tree, 7, 15))
}