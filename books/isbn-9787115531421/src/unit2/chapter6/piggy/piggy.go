package main

import (
	"fmt"
	"math/rand"
)

const unitCent = 0.05
var start = 0.0
const goal = 20.0

func main() {
	fmt.Printf("Now %.2f USD in the piggy bank.\n\n", start)
	putMoney()
}

func putMoney() {
	for {
		if start > goal {
			fmt.Printf("\nNow the piggy bank has %.2f USD, greater than %.2f USD.\n", start, goal)
			break
		}

		var i = rand.Intn(5) + 1
		if i == 1 || i == 2 || i == 5 {
			var increment = unitCent * float64(i)
			fmt.Printf("Put %.2f USD into the piggy bank.\n", increment)
			start += increment
			fmt.Printf("Now %.2f USD in the piggy bank.\n", start)
		}
	}
}