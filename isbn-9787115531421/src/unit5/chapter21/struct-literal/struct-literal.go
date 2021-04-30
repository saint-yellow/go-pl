package main

import "fmt"

func main() {
	type location struct {
		latitude, longitude float64
	}

	opportunity := location {latitude: -1.9462, longitude: 354.4734}
	fmt.Println("opportunity: ", opportunity)

	insight := location{latitude: 4.5, longitude: 135.9}
	fmt.Println("insight: ", insight)

	spirit := location{-14.5684, 175.472636}
	fmt.Println("spirit: ", spirit)

	curiosity := location{-4.5895, 137.4417}
	fmt.Printf("curiosity: %v\n", curiosity)
	fmt.Printf("curiosity: %+v\n", curiosity)
}