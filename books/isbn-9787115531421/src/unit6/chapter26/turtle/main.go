package main

import (
	"fmt"
	"math/rand"
)

func main() {
	t := turtle {
		x: 0, 
		y: 0,
	}

	fmt.Println("Step 0:", t)

	for i := 1; i <= 100; i++ {
		switch rand.Intn(5) {
			case 0:
				continue
			case 1:
				t.up()
			case 2:
				t.right()
			case 3:
				t.down()
			case 4:
				t.left()
		}
		fmt.Printf("Step%d: %v\n", i, t)
	}
}