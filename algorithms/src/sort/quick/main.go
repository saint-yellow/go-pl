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
	quickSort(numbers[:], 0, len(numbers)-1)
	fmt.Println("after:", numbers)
}

func quickSort(numbers []int, leftIndex, rightIndex int) {
	if rightIndex - leftIndex <= 0 {
		return
	}

	pivotPosition := partition(numbers, leftIndex, rightIndex)
	quickSort(numbers, leftIndex, pivotPosition-1)
	quickSort(numbers, pivotPosition+1, rightIndex)

}

func partition(numbers []int, leftPointer, rightPointer int) int {
	pivotPosition := rightPointer
	pivot := numbers[pivotPosition]

	rightPointer -= 1

	for {
		for numbers[leftPointer] < pivot {
			leftPointer += 1
		}
		for numbers[rightPointer] > pivot {
			rightPointer -= 1
		}

		if leftPointer >= rightPointer {
			break
		} else {
			swap(numbers, leftPointer, rightPointer)
		}
	}

	swap(numbers, leftPointer, pivotPosition)

	return leftPointer
}

func swap(numbers []int, pointer1, pointer2 int) {
	numbers[pointer1], numbers[pointer2] = numbers[pointer2], numbers[pointer1]
}
