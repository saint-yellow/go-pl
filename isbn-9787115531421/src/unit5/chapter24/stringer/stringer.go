package main

import "fmt"

type location struct {
	latitude, longitude float64
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.latitude, l.longitude)
}

func main() {
	curiosity := location {
		latitude: -4.5895,
		longitude: 137.4417,
	}

	fmt.Println(curiosity)
}