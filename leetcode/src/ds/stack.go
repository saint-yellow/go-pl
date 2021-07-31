package ds

type SinglyLinkedStack struct {
	sentinel *SinglyLinkedNode
}

func (s *SinglyLinkedStack) Empty() bool {
	return s.sentinel.Next == nil
}

func (s *SinglyLinkedStack) Push(x int) {
	pointer := s.sentinel
	for pointer.Next != nil {
		pointer = pointer.Next
	}
	pointer.Next = &SinglyLinkedNode{
		Val: x,
	}
}

func (s *SinglyLinkedStack) Top() int {
	pointer := s.sentinel
	for pointer.Next != nil {
		pointer = pointer.Next
	}
	return pointer.Val
}

func (s *SinglyLinkedStack) Pop() int {
	pointer := s.sentinel
	for pointer.Next != nil && pointer.Next.Next != nil {
		pointer = pointer.Next
	}
	x := pointer.Next.Val
	pointer.Next = nil
	return x
}

