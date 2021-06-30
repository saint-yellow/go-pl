package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

func main() {
	url := os.Args[1]
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("getting %s: %v\n", url, err)
		os.Exit(1)
	}
	doc, err := html.Parse(response.Body)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("parsing HTML: %s\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}
	

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}