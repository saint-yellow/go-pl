package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numbers := [30]int{}
	rand.Seed(time.Now().Unix())
	for i := range numbers {
		numbers[i] = rand.Intn(100)
	}
	fmt.Println("before:", numbers)
	selectionSort(numbers[:])
	fmt.Println("after:", numbers)
}

func selectionSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		lowestNumberIndex := i
		for j := i+1; j < len(numbers); j++ {
			if numbers[j] < numbers[lowestNumberIndex] {
				lowestNumberIndex = j
			}
		}

		if lowestNumberIndex != i {
			numbers[i], numbers[lowestNumberIndex] = numbers[lowestNumberIndex], numbers[i]
		}
	}
}