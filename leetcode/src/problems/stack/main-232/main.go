// LeetCode 主站 Problem Nr. 232: 用栈实现队列

package main

// 用栈实现的队列
type MyQueue struct {
	inputStack, outputStack []int
}

// 构造函数
func Constructor() MyQueue {
	return MyQueue{
		inputStack: make([]int, 0),
		outputStack: make([]int, 0),
	}
}

// 将元素 x 推到队列的末尾
func (q *MyQueue) Push(x int) {
	q.inputStack = append(q.inputStack, x)
}

// 从队列的开头移除并返回元素
func (q *MyQueue) Pop() int {
	return q.head(true)
}

// 返回队列开头的元素
func (q *MyQueue) Peek() int {
	return q.head(false)
}

// 如果队列为空，返回 true；否则，返回 false
func (q *MyQueue) Empty() bool {
	return (len(q.inputStack) == 0 && len(q.outputStack) == 0)
}

// 将inputStack的元素搬移到outputStack
func (q *MyQueue) fromInToOut() {
	for len(q.inputStack) > 0 {
		n := len(q.inputStack)
		q.outputStack = append(q.outputStack, q.inputStack[n-1])
		q.inputStack = q.inputStack[:n-1]
	}
}

// 返回队列开头的元素，当popHead=true时，删除该元素
func (q *MyQueue) head(popHead bool) int {
	if len(q.outputStack) == 0 {
		q.fromInToOut()
	}

	x := q.outputStack[len(q.outputStack) - 1]
	if popHead {
		q.outputStack = q.outputStack[:len(q.outputStack)-1]
	}
	return x
}

func main() {
	q := MyQueue{}
	q.Push(1)
	q.Push(2)
	q.Peek()
	q.Pop()
	q.Empty()
}