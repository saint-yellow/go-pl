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
	insertionSort(numbers[:])
	fmt.Println("after:", numbers)
}


func insertionSort(numbers []int) {
	for i := 1; i < len(numbers); i++ {
		position := i
		tempValue := numbers[position]

		for position > 0 && numbers[position-1] > tempValue {
			numbers[position] = numbers[position-1]
			position -= 1
		}

		numbers[position] = tempValue
	}
}