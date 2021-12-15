package main

import (
	"fmt"
	"math"
)

const (
	frequency = 10
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	roots := make(chan float64)

	go func() {
		for x := 0; x <= frequency; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	
	// squares <- naturals
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// roots <- naturals
	go func() {
		for x := range naturals {
			roots <- math.Sqrt(float64(x))
		}
		close(roots) 
	}()

	// for x := range squares {
	// 	fmt.Println(x)
	// }

	for x := range roots {
		fmt.Println(x)
	}
}