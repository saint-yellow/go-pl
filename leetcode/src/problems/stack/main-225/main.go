// LeetCode 主站 Problem Nr. 225: 用队列实现栈

package main

type MyStack struct {
	queue1, queue2 []int
}

func Constructor() MyStack {
	return MyStack{
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}

func (s *MyStack) Push(x int) {
    s.queue2 = append(s.queue2, x)
    for len(s.queue1) > 0 {
        s.queue2 = append(s.queue2, s.queue1[0])
        s.queue1 = s.queue1[1:]
    }
    s.queue1, s.queue2 = s.queue2, s.queue1
}

func (s *MyStack) Pop() int {
	return s.head(true)
}

func (s *MyStack) Top() int {
	return s.head(false)
}

func (s *MyStack) head(popHead bool) int {
	v := s.queue1[0]
	if popHead {
		s.queue1 = s.queue1[1:]
	}
	return v
}

func (s *MyStack) Empty() bool {
	return len(s.queue1) == 0
}