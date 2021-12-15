package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)



func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercise-05.03: %v\n", err)
		os.Exit(1)
	}

	printContent(doc)
}

func printContent(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Printf("%s\n", n.Data)
	}

	// for c := n.FirstChild; c != nil; c = n.NextSibling {
	// 	if c.Data == "script" || c.Data == "style" {
	// 		continue
	// 	}
	// 	printContent(c)
	// }
}