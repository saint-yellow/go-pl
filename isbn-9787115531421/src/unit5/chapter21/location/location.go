package main

import "fmt"


type location struct {
	latitude float64
	longitude float64
}

func main() {
	var spirit location
	spirit.latitude = -14.5684
	spirit.longitude = 175.472

	var opportunity location
	opportunity.latitude = -1.9462
	opportunity.longitude = 354.4734

	fmt.Println(spirit, opportunity)
}