// LeetCode 主站 Problem Nr. 1670: 设计前中后队列

package main

type FrontMiddleBackQueue struct {
	array []int
}


func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		array: make([]int, 0),
	}
}


func (this *FrontMiddleBackQueue) PushFront(val int)  {
	this.array = append([]int{val}, this.array...)
}


func (this *FrontMiddleBackQueue) PushMiddle(val int)  {
	m := len(this.array) / 2
	this.array = append(this.array[:m], append([]int{val}, this.array[m:]...)...)
}


func (this *FrontMiddleBackQueue) PushBack(val int)  {
	this.array = append(this.array, val)
}


func (this *FrontMiddleBackQueue) PopFront() int {
	if len(this.array) == 0 {
		return -1
	}

	val := this.array[0]
	this.array = this.array[1:]
	return val
}


func (this *FrontMiddleBackQueue) PopMiddle() int {
	if len(this.array) == 0 {
		return -1
	}

	m := (len(this.array) - 1) / 2
	val := this.array[m]
	this.array = append(this.array[:m], this.array[m+1:]...)
	return val
}


func (this *FrontMiddleBackQueue) PopBack() int {
	if len(this.array) == 0 {
		return -1
	}

	val := this.array[len(this.array)-1]
	this.array = this.array[:len(this.array)-1]
	return val
}