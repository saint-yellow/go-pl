package main

import (
	"fmt"
	"os"
	"text/tabwriter"
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

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _,t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}