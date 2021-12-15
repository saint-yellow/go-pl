package main

import (
	"fmt"
	"net/http"
	"os"

	"strings"

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
	attrs := make([]string, 0)
	for _, a := range n.Attr {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, a.Key, a.Val))
	}

	if n.FirstChild == nil {
		fmt.Printf("%*s<%s/ %s>\n", depth*2, "", n.Data, strings.Join(attrs, " "))
	} else {
		fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Data, strings.Join(attrs, " "))
	}

	depth++
}

func endElement(n *html.Node) {
	depth--

	if n.FirstChild == nil {
		return
	}
	fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
}