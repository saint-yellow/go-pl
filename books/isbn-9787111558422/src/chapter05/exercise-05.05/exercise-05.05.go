package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := parseHTMLDoc(os.Args[1])
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	words, images := countWordsAndImages(doc)
	fmt.Printf("Page: %s, Words: %d, Images: %d\n", os.Args[1], words, images)
}

func parseHTMLDoc(url string) (*html.Node, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(response.Body)
	defer response.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return nil, err
	}
	return doc, nil
}

func countWordsAndImages(n *html.Node) (words, images int) {
	u := make([]*html.Node, 0)
	u = append(u, n)
	for len(u) > 0 {
		n = u[len(u)-1]
		u = u[:len(u)-1]
		switch n.Type {
		case html.TextNode:
			words += wordCount(n.Data)
		case html.ElementNode:
			if n.Data == "img" {
				images++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			u = append(u, c)
		}
	}
	return
}

func wordCount(s string) (count int) {
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		count++
	}
	return
}
