// 剑指Offer Problem Nr. 10-I: 斐波那契数列

package main

import "fmt"

func fib(n int) int {
	return method1(n)
}

func method1(n int) int {
	if n <= 1 {
		return n
	}

	const mod int = 1e9 + 7
	p, q, r := 0, 0 ,1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (p+q) % mod
	}
	return r
}

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(95))
}