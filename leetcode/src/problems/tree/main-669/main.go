// LeetCode 主站 Problem Nr. 669: 修建二叉搜索树

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	return method1(root, low, high)
}

func method1(root *TreeNode, low, high int) *TreeNode {
	// ...
	return nil
}

func main() {
	trimBST(nil, 0, 1)
}