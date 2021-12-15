package main

import (
	"flag"
	"isbn-9787111558422/src/chapter04/exercise-04.12/xkcd"
	"time"
	"log"
	"fmt"
)



func main() {
	comicNo := flag.Int("n", int(xkcd.LatestComic), "Comic number to fetch (default latest)")
	clientTimeout := flag.Int64("t", int64(xkcd.DefaultClientTimeout.Seconds()), "Client timeout in seconds")
	downloadImage := flag.String("d", "", "Download image to specific directory")
	outputType := flag.String("o", "text", "Print output in format: text/json")
	flag.Parse()
	
	xc := xkcd.NewXKCDClient()
	xc.SetTimeout(time.Duration(*clientTimeout) * time.Second)
	comic, err := xc.Fetch(xkcd.ComicNumber(*comicNo), *downloadImage)
	if err != nil {
		log.Println(err)
	}
	if *outputType == "json" {
		fmt.Println(comic.JSON())
	} else {
		fmt.Println(comic.PrettyString())
	}
}