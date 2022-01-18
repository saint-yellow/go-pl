// LeetCode 主站 Problem Nr. 239: 滑动窗口最大值

package main

import "fmt"


type queue struct {
	data []int
}
func (q *queue) pop(value int) {
	if !q.empty() && value == q.front() {
		q.data = q.data[1:]
	}
}
func (q *queue) push(value int) {
	for !q.empty() && value > q.back() {
		q.data = q.data[:len(q.data)-1]
	}
	q.data = append(q.data, value)
}

// 返回队头元素
func (q *queue) front() int { return q.data[0] }

// 返回队尾元素
func (q *queue) back() int { return q.data[len(q.data)-1] }


func (q *queue) empty() bool { return len(q.data) == 0 }

func maxSlidingWindow(numbers []int, k int) []int {
	q := queue{
		data: make([]int, len(numbers)-k+1),
	}
	result := make([]int, 0)
	for i := 0; i < k; i++ {
		q.push(numbers[i])
	}
	result = append(result, q.front())
	for i := k; i < len(numbers); i++ {
		q.pop(numbers[i-k])
		q.push(numbers[i])
		result = append(result, q.front())
	}
	return result
}

func main() {
	numbers := []int{1,3,-1,-3,5,3,6,7}
	k := 3

	fmt.Println(maxSlidingWindow(numbers, k))
}