package main

import (
	"fmt"
)

type location struct {
	latitude, longitude float64
}

type coordinate struct {
	d, m, s float64
	h rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(latitude, longitude coordinate) location {
	return location{latitude: latitude.decimal(), longitude: longitude.decimal()}
}

func main() {
	latitude := coordinate{4, 35, 22.2, 'S'}
	longitude := coordinate{137, 26, 30.12, 'E'}
	curiosity := newLocation(latitude, longitude)
	fmt.Println(curiosity)
}