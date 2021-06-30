package main

import "fmt"

func main() {
	var curiosity struct {
		latitude float64
		longitude float64
	}

	curiosity.latitude = -4.5895
	curiosity.longitude = 137.4417

	fmt.Println(curiosity.latitude, curiosity.longitude)
	fmt.Println(curiosity)
}