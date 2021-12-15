package main

import (
	"html/template"
	"io"
	"log"
	"time"
)

type Track struct {
	Title, Artist, Album string
	Year int
	Length time.Duration
}

func length(s string) time.Duration {
	if d, err := time.ParseDuration(s); err != nil {
		panic(s)
	} else {
		return d
	}
}

func printTracks(writer io.Writer, tracks []*Track) {
	tracksTemplate := template.Must(template.ParseFiles("index.html"))
	if err := tracksTemplate.Execute(writer, tracks); err != nil {
		log.Fatal(err)
	}

}