// LeetCode 主站 Problem Nr. 59: 螺旋矩阵II

package main

import "fmt"

func generateMatrix(n int) [][]int {
	return method1(n)
}

func method1(n int) [][]int {
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1
	tar := n * n
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for num <= tar {
		// 在顶部从左向右填数
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}

		// 顶部下沉
		top++

		// 在右侧从上到下填数
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}

		// 右侧左移
		right--

		// 在底部从右到左填数
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}

		// 底部上升
		bottom--

		// 在左侧从下到上填数
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}

		// 左侧右移
		left++
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(4))
}