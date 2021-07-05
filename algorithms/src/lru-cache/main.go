package main

import "fmt"

type SinglyLinkedList struct {
	value int
	next *SinglyLinkedList
}

type LRUCache struct {
	head *SinglyLinkedList
	capacity int
}

func (lc LRUCache) length() int {
	length := 0
	fakeHead := lc.head
	for fakeHead != nil {
		length += 1
		fakeHead = fakeHead.next
	}
	return length
}

func (lc LRUCache) get(value int) (int, error) {
	fakeHead := &SinglyLinkedList{
		value: -1,
		next: lc.head,
	}

	for fakeHead.next != nil {
		if fakeHead.next.value == value {
			return value, nil
		}
		fakeHead = fakeHead.next
	}

	return -1, fmt.Errorf("value %d not found in cache", value)
}

func (lc LRUCache) set(value int) bool {
	_, err := lc.get(value)
	if err != nil {
		// add new value to the cache
		if lc.length() < lc.capacity {
			node := &SinglyLinkedList{
				value: value,
			}
			node.next = lc.head
			return true
		}

		if lc.length() == lc.capacity {
			fakeHead := &SinglyLinkedList{
				value: -1,
				next: lc.head,
			}

			for fakeHead.next.next != nil {
				fakeHead = fakeHead.next
			}
			fakeHead.next = nil
			node := &SinglyLinkedList{
				value: value,
			}
			node.next = lc.head
			return true
		}

	} else {
		return false
	}

	return false
}