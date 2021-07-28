package ds

type BinaryNode struct {
	Val int
	Left *BinaryNode
	Right *BinaryNode
}

type NAryNode struct {
	Val int
	Children []*NAryNode
}