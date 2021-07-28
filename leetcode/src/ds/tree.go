package ds

type BinaryNode struct {
	Val int
	Left *BinaryNode
	Right *BinaryNode
}

func (root *BinaryNode) preOrderTraversal() []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	result = append(result, root.Val)
	result = append(result, root.Left.preOrderTraversal()...)
	result = append(result, root.Right.preOrderTraversal()...)
	return result
}

func (root *BinaryNode) inOrderTraversal() []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	result = append(result, root.Left.inOrderTraversal()...)
	result = append(result, root.Val)
	result = append(result, root.Right.inOrderTraversal()...)
	return result
}

func (root *BinaryNode) postOrderTraversal() []int {
	result := make([]int, 0)
	
	if root == nil {
		return result
	}

	result = append(result, root.Left.postOrderTraversal()...)
	result = append(result, root.Right.postOrderTraversal()...)
	result = append(result, root.Val)
	return result
}

// binary tree's depth-first traversal
func (root *BinaryNode) DFT(rootOrder int) []int {
	switch rootOrder {
		case 0:
			return root.preOrderTraversal()
		case 1:
			return root.inOrderTraversal()
		default:
			return root.postOrderTraversal()
	}
}

func (root *BinaryNode) OneDimensionalBFT() []int {
	queue := make([]*BinaryNode, 0)
	if root != nil {
		queue = append(queue, root)
	}

	result := make([]int, 0)
	for len(queue) > 0 {
		size := len(queue)
		currentLevel := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, currentLevel...)
	}

	return result
}

func (root *BinaryNode) TwoDimensionalBFT() [][]int {
	queue := make([]*BinaryNode, 0)
	if root != nil {
		queue = append(queue, root)
	}

	result := make([][]int, 0)
	for len(queue) > 0 {
		size := len(queue)
		currentLevel := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, currentLevel)
	}

	return result
}

func BuildBinaryTree(values []int) *BinaryNode {
	return &BinaryNode{Val: -1}
}

type NAryNode struct {
	Val int
	Children []*NAryNode
}

func (root *NAryNode) OneDimensionalBFT() []int {
	queue := make([]*NAryNode, 0)

	if root != nil {
		queue = append(queue, root)
	}

	result := make([]int, 0)
	
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Children != nil {
				queue = append(queue, node.Children...)
			}
		}
		result = append(result, level...)
	}

	return result
}

func (root *NAryNode) TwoDimensionalBFT() [][]int {
	queue := make([]*NAryNode, 0)

	if root != nil {
		queue = append(queue, root)
	}

	result := make([][]int, 0)
	
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Children != nil {
				queue = append(queue, node.Children...)
			}
		}
		result = append(result, level)
	}

	return result
}