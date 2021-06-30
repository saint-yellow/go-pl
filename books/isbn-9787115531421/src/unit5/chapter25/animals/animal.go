package main

import (
	"fmt"
	"math/rand"
)



type animal interface {
	String() string
	move() string
	eat() string
}

func step(a animal) {
	switch rand.Intn(2) {
		case 0:
			fmt.Printf("%v %v.\n", a, a.move())
		default:
			fmt.Printf("%v eats the %v.\n", a, a.eat())
	}
}