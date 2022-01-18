// LeetCode 主站 Problem Nr. 291: 存在重复元素 II

package main

import "fmt"

func containsNearbyDuplicate(numbers []int, k int) bool {
	return method1(numbers, k)
}

func method1(numbers []int, k int) bool {
	d := make(map[int]int)
	for i, v := range numbers {
		if n, ok := d[v]; ok && i - n <= k {
			return true
		}
		d[v] = i
	}
	return false
}

func method2(numbers []int, k int) bool {
	d := map[int]int{}
	for i, v := range numbers {
		if i > k {
			delete(d, numbers[i-k-1])
		}
		if _, ok := d[v]; ok {
			return true
		}
		d[v] = -1
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1,2,3,1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1,0,1,1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1,2,3,1,2,3}, 2))

	fmt.Println(method2([]int{1,2,3,1,2,3}, 2))
}
