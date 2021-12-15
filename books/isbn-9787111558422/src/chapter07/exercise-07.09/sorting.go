package main

import (
	"log"
	"net/http"
)

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func main () {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	// sort.Sort(CustomSort{tracks, func(x, y *Track) bool {
	// 	if x.Title != y.Title {
	// 		return x.Title < y.Title
	// 	}

	// 	if x.Artist != y.Artist {
	// 		return x.Artist < y.Artist
	// 	}

	// 	if x.Length != y.Length {
	// 		return x.Length < y.Length
	// 	}

	// 	return false
	// }})
	printTracks(w, tracks)
}