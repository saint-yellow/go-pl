package main

import (
	"fmt"
	"strings"
)

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	go sourceGopher(c0)
	go splitterGopher(c0, c1)
	printerGopher(c1)
}

func sourceGopher(downstream chan string) {
	content := []string { "the quick brown fox", "jumps over", "the lazy dog" }
	for _, v := range content {
		downstream <- v
	}
	close(downstream)
}

func splitterGopher(upstream, downstream chan string) {
	for phrase := range upstream {
		for _, word := range strings.Fields(phrase) {
			downstream <- word
		}
	}
	close(downstream)
}

func printerGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}