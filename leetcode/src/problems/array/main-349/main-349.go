// LeetCode 主站 Problem Nr. 349: 两个数组的交集

package main

func intersection(numbers1, numbers2 []int) []int {
	return method1(numbers1, numbers2)
}

func method1(numbers1, numbers2 []int) []int {
	set := make(map[int]int)
	result := make([]int, 0)

	for _, n := range numbers1 {
		if _, ok := set[n]; !ok {
			set[n] = 1
		}
	}

	for _, n := range numbers2 {
		if _, ok := set[n]; ok {
			result = append(result, n)
			delete(set, n)
		}
	}
	return result
}

func method2(numbers1, numbers2 []int) []int {
	set := make(map[int]int)
	result := make([]int, 0)

	for _, n := range numbers1 {
		if _, ok := set[n]; !ok {
			set[n] = 1
		}
	}

	for _, n := range numbers2 {
		if count, ok := set[n]; ok && count > 0 {
			result = append(result, n)
			set[n]--
		}
	}
	return result
}
