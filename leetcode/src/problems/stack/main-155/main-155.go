// LeetCode 主站 Problem Nr. 155: 最小栈

package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack []int
	minStack []int
}


func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		minStack: []int{math.MaxInt64},
	}
}


func (ms *MinStack) Push(val int)  {
	ms.stack = append(ms.stack, val)
	minValue := ms.minStack[len(ms.minStack)-1]
	if val < minValue {
		ms.minStack = append(ms.minStack, val)
	} else {
		ms.minStack = append(ms.minStack, minValue)
	}

}


func (ms *MinStack) Pop()  {
	ms.stack = ms.stack[:len(ms.stack)-1]
	ms.minStack = ms.minStack[:len(ms.minStack)-1]
}


func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}


func (ms *MinStack) GetMin() int {
	return ms.minStack[len(ms.minStack)-1]
}

func (ms *MinStack) Show() {
	fmt.Println(ms)
}


func main() {
	stack := Constructor()
	stack.Show()

	stack.Push(-2)
	stack.Show()

	stack.Push(0)
	stack.Show()

	stack.Push(-3)
	stack.Show()

	fmt.Println("min:", stack.GetMin())

	stack.Pop()
	stack.Show()

	fmt.Println("top:", stack.Top())
	stack.Show()

	fmt.Println("min:", stack.GetMin())
}