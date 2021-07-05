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
	bubbleSort(numbers[:])
	fmt.Println("after:", numbers)
}

func bubbleSort(numbers []int) {
	unsortedUntilIndex := len(numbers) - 1
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < unsortedUntilIndex; i++ {
			if numbers[i] > numbers[i+1] {
				sorted = false
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			}
		}
		unsortedUntilIndex -= 1
	}

}