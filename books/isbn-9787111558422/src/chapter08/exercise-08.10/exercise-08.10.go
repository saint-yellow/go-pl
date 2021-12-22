package main

import (
	"fmt"
	"isbn-9787111558422/src/chapter08/exercise-08.10/links"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl(url string, cancelled chan struct{}) []string {
    fmt.Println(url)
	tokens <- struct{}{}
    list, err := links.Extract(url, cancelled)
	<- tokens
    if err != nil {
        log.Print(err)
    }
    return list
}

func main() {
	workList := make(chan []string)
	var n int

	n++
	go func() {
		workList <- os.Args[1:]
	}()

	cancelled := make(chan struct{})

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-workList
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					workList <- crawl(link, cancelled)
				}(link)
			}
		}
	}
}