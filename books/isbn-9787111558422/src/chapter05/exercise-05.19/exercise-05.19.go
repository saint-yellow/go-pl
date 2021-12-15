package main

import "fmt"

func main() {
	fmt.Println(recoverTest(10))
}

func recoverTest(x int) (result int) {
	defer func() {
		recover()
		result = x
	}()

	panic(x)
}