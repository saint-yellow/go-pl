package ds

type SinglyLinkedNode struct {
	Value int
	Next *SinglyLinkedNode
}

type DoublyLinkedNode struct {
	Value int
	Prev *DoublyLinkedNode
	Next *DoublyLinkedNode
}