package main

import (
	"fmt"
	"isbn-9787115531421/src/unit5/chapter22/distance/world"
	"isbn-9787115531421/src/unit5/chapter22/landing/coordinate"
	"isbn-9787115531421/src/unit5/chapter22/landing/location"
)

var (
	mars = world.New(3389.5)
	earth = world.New(6371)
)


func main() {
	spirit := location.New("Spirit", coordinate.New(14, 34, 6.2, 'S'), coordinate.New(175, 28, 21.5, 'E'))
	opportunity := location.New("Opportunity", coordinate.New(1, 56, 46.3, 'S'), coordinate.New(354, 28, 24.2, 'E'))
	curiosity := location.New("Curiosity", coordinate.New(4, 35, 22.2, 'S'), coordinate.New(137, 26, 30.12, 'E'))
	insight := location.New("Insight", coordinate.New(4, 30, 0.0, 'N'), coordinate.New(135, 54, 0, 'E'))

	fmt.Printf("Spirit to Opportunity %.2f km\n", mars.Distance(spirit, opportunity))
	fmt.Printf("Spirit to Curiosity %.2f km\n", mars.Distance(spirit, curiosity))
	fmt.Printf("Spirit to Insight %.2f km\n", mars.Distance(spirit, insight))
	fmt.Printf("Opportunity to Curiosity %.2f km\n", mars.Distance(opportunity, curiosity))
	fmt.Printf("Opportunity to Insight %.2f km\n", mars.Distance(opportunity, insight))
	fmt.Printf("Curiosity to Insight %.2f km\n", mars.Distance(curiosity, insight))

	london := location.New("London", coordinate.New(51, 30, 0, 'N'), coordinate.New(0, 8, 0, 'W'))
	paris := location.New("Paris", coordinate.New(48, 51, 0, 'N'), coordinate.New(2, 21, 0, 'E'))

	fmt.Printf("London to Paris %.2f km\n", earth.Distance(london, paris))
}