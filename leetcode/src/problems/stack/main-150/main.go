// LeetCode 主站 Problem Nr. 150: 逆波兰表达式求值

package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	return method1(tokens)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return b - a
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return b / a
}

var computations = map[string]func(int, int)int {
	"+": add,
	"-": subtract,
	"*": multiply,
	"/": divide,
} 

func method1(tokens []string) int {
	stack := make([]int, 0)
	for _, t := range tokens {
		if digit, err := strconv.Atoi(t); err == nil {
			stack = append(stack, digit)
		} else {
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			c := computations[t](a, b)
			stack = append(stack, c)

		}
	}
	return stack[0]
}

func main() {
	fmt.Println(evalRPN([]string{"2","1","+","3","*"}))
	fmt.Println(evalRPN([]string{"4","13","5","/","+"}))
	fmt.Println(evalRPN([]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}))
}