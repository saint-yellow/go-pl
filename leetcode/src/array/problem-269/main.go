package main

import "math"

func majorElement(numbers []int) int {
	length := len(numbers)
	if length == 1 {
		return numbers[0]
	}


	half := int(math.Floor(float64(length/2)))
	hash := make(map[int]int, 0)

	for _, number := range numbers {
		if _, ok := hash[number]; ok {
			hash[number] += 1
			if hash[number] > half {
				return number
			}
		} else {
			hash[number] = 1
		}
	}

	return -1
}