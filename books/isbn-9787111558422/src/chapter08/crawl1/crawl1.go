package main

import (
	"fmt"
	"isbn-9787111558422/src/chapter05/links"
	"log"
	"os"
)

func crawl(url string) []string {
    fmt.Println(url)
    list, err := links.Extract(url)
    if err != nil {
        log.Print(err)
    }
    return list
}

func main() {
	workList := make(chan []string)

	go func() {
		workList <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}
}