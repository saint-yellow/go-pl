package main

import "fmt"

func main() {
	s := []string{}
	lastCapacity := cap(s)
	for i := 0; i < 10000; i++ {
		s = append(s, "An element")
		if (cap(s) != lastCapacity) {
			fmt.Println(cap(s))
			lastCapacity = cap(s)
		}
	}

	fmt.Println()

	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s[:10])
}