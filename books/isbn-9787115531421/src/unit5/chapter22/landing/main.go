package main

import (
	"fmt"
	"isbn-9787115531421/src/unit5/chapter22/landing/coordinate"
	"isbn-9787115531421/src/unit5/chapter22/landing/location"
)


func main() {
	

	spirit := location.New("Spirit", coordinate.New(14, 34, 6.2, 'S'), coordinate.New(175, 28, 21.5, 'E'))
	opportunity := location.New("Opportunity", coordinate.New(1, 56, 46.3, 'S'), coordinate.New(354, 28, 24.2, 'E'))
	curiosity := location.New("Curiosity", coordinate.New(4, 35, 22.2, 'S'), coordinate.New(137, 26, 30.12, 'E'))
	insight := location.New("Insight", coordinate.New(4, 30, 0.0, 'N'), coordinate.New(135, 54, 0, 'E'))

	fmt.Println(spirit)
	fmt.Println(opportunity)
	fmt.Println(curiosity)
	fmt.Println(insight)
}