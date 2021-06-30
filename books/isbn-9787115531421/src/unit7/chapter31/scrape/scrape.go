package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type visitMap struct {
	mu sync.Mutex
	visited map[string]int
}

func (v *visitMap) visitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()

	count := v.visited[url]
	count++
	v.visited[url] = count

	return count
}

func main() {
	vm := visitMap {
		mu:sync.Mutex{}, 
		visited: make(map[string]int), 
	}
	for i := 0; i < 30; i++ {
		var url string
		switch rand.Int31n(2) {
			case 0:
				url = "https://python.org"
			default:
				url = "https://go.dev"
		}
		go vm.visitLink(url)
	}

	fmt.Printf("%+v\n", vm.visited)
}