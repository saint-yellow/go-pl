package main

import (
	"fmt"
	"math/rand"
)

const unitCent = 5
var start = 0
const goal = 20000

func main() {
	fmt.Printf("Now %.2f USD in the piggy bank.\n\n", float64(start)/100)
	putMoney(start, goal)
}

func putMoney(start int, goal int) {
	for {
		if start > goal {
			fmt.Printf("\nNow the piggy bank has %.2f USD, greater than %.2f USD.\n", float64(start)/100, float64(goal)/100)
			break
		}

		var i = rand.Intn(5) + 1
		if i == 1 || i == 2 || i == 5 {
			var increment = unitCent * i
			fmt.Printf("Put %d cents into the piggy bank.\n", increment)
			start += increment
			fmt.Printf("Now %.2f USD in the piggy bank.\n", float64(start)/100)
		}
	}
}