package main

import "fmt"

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printerGopher(c1)
}

func sourceGopher(downstream chan string) {
	content := []string {
		"a", "b", "b", "c", "d", "d", "e", "f", "f", "g",
	}
	for _, v := range content {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	previous := ""
	for v := range upstream {
		if v != previous {
			downstream <- v
			previous = v
		}
	}
	close(downstream)
}

func printerGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}