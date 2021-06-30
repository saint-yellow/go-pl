package main

import "fmt"

func main() {
	type location struct {
		latitude, longitude float64
	}

	bradbury := location{latitude: -4.5895, longitude: 137.4417}
	curiosity := bradbury

	curiosity.longitude += 0.0106

	fmt.Println(bradbury, curiosity)
}