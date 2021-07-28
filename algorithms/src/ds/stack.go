package ds

type SinglyLinkedStack struct {
	size int
	sentinel *SinglyLinkedNode
}

func NewSinglyLinkedStack() SinglyLinkedStack {
	return SinglyLinkedStack{
		size: 0,
		sentinel: &SinglyLinkedNode{
			Value: -1,
			Next: nil,
		},
	}
}

func (stack SinglyLinkedStack) Top() *SinglyLinkedNode {
	return stack.sentinel.Next
}

func (stack SinglyLinkedStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack SinglyLinkedStack) Size() int {
	return stack.size
}

func (stack SinglyLinkedStack) Push(value int) {
	node := &SinglyLinkedNode{
		Value: value,
		Next: stack.sentinel.Next,
	}
	stack.sentinel.Next = node
	stack.size += 1
}

func (stack SinglyLinkedStack) Pop() *SinglyLinkedNode {
	if stack.size == 0 {
		return nil
	}

	node := stack.sentinel.Next
	stack.sentinel.Next = node.Next
	node.Next = nil
	stack.size -= 1

	return node
}