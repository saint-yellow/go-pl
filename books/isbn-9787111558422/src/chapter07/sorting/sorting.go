package main

import (
	"fmt"
	"sort"
)

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func main () {
	sort.Sort(SortByArtist(tracks))
	printTracks(tracks)

	fmt.Println()

	// todo: reverse
	// sort.Reverse(SortByArtist(tracks))
	// printTracks(tracks)

	fmt.Println()

	sort.Sort(CustomSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}

		if x.Artist != y.Artist {
			return x.Artist < y.Artist
		}

		if x.Length != y.Length {
			return x.Length < y.Length
		}

		return false
	}})
	printTracks(tracks)
}