package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	for i := 0; i < 5; i++ {
		fmt.Println(f()) // 1, 4, 9, 16, 25
	}

	fmt.Println(f()) // 36
	
}