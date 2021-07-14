package ds

type SinglyLinkedNode struct {
	Val int
	Next *SinglyLinkedNode
}

func (node *SinglyLinkedNode) ToList() []int {
	result := make([]int, 0)
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}