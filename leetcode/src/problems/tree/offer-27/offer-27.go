// 剑指 Offer Problem Nr. 27: 二叉树的镜像

package main

import "github.com/saint-yellow/go-pl/leetcode/src/ds"

type TreeNode = ds.BinaryNode

func mirrorTree(root *TreeNode) *TreeNode {
	return method1(root)
}

func method1(root *TreeNode) *TreeNode {
	if root == nil {
        return nil
    }

    root.Left, root.Right = root.Right, root.Left
    if root.Left != nil {
        mirrorTree(root.Left)
    }
    if root.Right != nil {
        mirrorTree(root.Right)
    }

    return root
}

func method2(root *TreeNode) *TreeNode {
	queue := make([]*TreeNode, 0)

	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}