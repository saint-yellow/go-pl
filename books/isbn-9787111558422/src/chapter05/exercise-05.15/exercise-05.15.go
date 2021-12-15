package main

import "fmt"

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("at least one number is required")
	}

	max := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] > max {
			max = vals[i]
		}
	}
	return max, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("at least one number is required")
	}

	min := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] < min {
			min = vals[i]
		}
	}
	return min, nil
}

func main() {
	fmt.Println(max(1,2,3,4))
	fmt.Println(min())
}