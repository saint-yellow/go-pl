package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercise-05.02: %v\n", err)
		os.Exit(1)
	}

	var result = map[string]int{}
	stat(&result, doc)
	fmt.Println(result)
}

func stat(result *map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		(*result)[n.Data]++
	}
	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	stat(result, c)
	// }

	if n.FirstChild != nil {
		stat(result, n.FirstChild)
	}

	if n.NextSibling != nil {
		stat(result, n.NextSibling)
	}
}