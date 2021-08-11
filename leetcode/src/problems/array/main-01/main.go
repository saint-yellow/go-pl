package main

func twoSum(numbers []int, target int) [2]int {
	return method1(numbers, target)
}

func method1(numbers []int, target int) [2]int {
	length := len(numbers)
	dict := make(map[int]int, 0)
	for i := 0; i < length; i++ {
		a := target - numbers[i]
		if _, ok := dict[numbers[i]]; ok {
			return [2]int{dict[numbers[i]], i}
		}
		dict[a] = i
	}
	return [2]int{-1, -1}
}