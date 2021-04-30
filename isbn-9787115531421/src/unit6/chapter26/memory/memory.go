package main

import "fmt"

func main() {
	// 定义一个变量
	answer := 42
	fmt.Println(answer)
	fmt.Println(&answer)

	// 定义一个指针
	address := &answer
	fmt.Println(*address)

	fmt.Println(*&answer)
}