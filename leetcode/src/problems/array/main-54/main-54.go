// LeetCode 主站 Problem Nr. 54: 螺旋矩阵

package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	return method1(matrix)
}

func method1(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])

	top, bottom := 0, m-1
	left, right := 0, n-1
	start, end := 0, m*n-1
	result := make([]int, m*n)
	for start <= end {
		for i := left; i <= right; i++ {
			result[start] = matrix[top][i]
			start++
		}

		top++
		if (top > bottom) {
			break
		}

		for i := top; i <= bottom; i++ {
			result[start] = matrix[i][right]
			start++
		}

		right--
		if (right < left) {
			break
		}


		for i := right; i >= left; i-- {
			// fmt.Printf("i=%d,start=%d\n", i, start)
			result[start] = matrix[bottom][i]
			start++
		}

		bottom--
		if (bottom < top) {
			break
		}

		for i := bottom; i >= top; i-- {
			result[start] = matrix[i][left]
			start++
		}

		left++
		if (left > right) {
			break
		}
	}

	return result
}

func main() {
	matrix := [][]int{
		{1,2,3,4},
		{12,13,14,5},
		{11,16,15,6},
		{10,9,8,7},
		{100,100,100,100},
	}
	fmt.Println(spiralOrder(matrix))
	fmt.Println(spiralOrder(make([][]int, 0)))
}