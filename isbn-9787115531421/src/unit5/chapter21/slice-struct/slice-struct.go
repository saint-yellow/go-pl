package main

import "fmt"

func main() {
	type location struct {
		name string
		latitude float64
		longitude float64
	}

	locations := []location {
		{name: "Brabury Landing", latitude: -4.5895, longitude:137.4417},
		{name: "Columbia Memorial Station", latitude: -14.5684, longitude: 175.472636},
		{name: "Challenger Memorial Station", latitude: -1.9462, longitude: 354.4734},
	}

	fmt.Println(locations)

	for _, location := range locations {
		fmt.Println(location)
	}
}