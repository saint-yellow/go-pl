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

func BuildSinglyLinkedList(values []int) *SinglyLinkedNode {
	sentinel := &SinglyLinkedNode{Val: -1}
	pointer := sentinel

	for _, v := range values {
		pointer.Next = &SinglyLinkedNode{Val: v}
		pointer = pointer.Next
	}

	return sentinel.Next
}


type DoublyLinkedNode struct {
	Val int
	Next, Previous *DoublyLinkedNode
}

func BuildDoublyLinkedList(values []int) *DoublyLinkedNode {
	sentinel := &DoublyLinkedNode{Val: -1}
	pointer := sentinel

	for _, v := range values {
		previous := pointer
		pointer.Next = &DoublyLinkedNode{Val: v}
		pointer.Previous = previous
	}
	
	return sentinel.Next
}