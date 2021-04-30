package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var result = 75
	for {
		var random = rand.Intn(100) + 1
		fmt.Printf("I guess: answer = %v? ", random)
		if random == result {
			fmt.Printf("%v = answer.\n", random)
			break
		} else if random > result {
			fmt.Printf("%v > answer.\n", random)
		} else {
			fmt.Printf("%v < answer.\n", random)
		}
	}
	fmt.Println("Done.")
}