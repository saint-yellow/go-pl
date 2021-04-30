package main

import (
	"fmt"
	"math"
)

type World struct {
	radius float64
}

type location struct {
	latitude, longitude float64
}

func (w World) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.latitude))
	s2, c2 := math.Sincos(rad(p2.latitude))
	clong := math.Cos(rad(p1.longitude - p2.longitude))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {
	mars := World{3389.5}

	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}
	distance := mars.distance(spirit, opportunity)

	fmt.Printf("%.2fkm\n", distance)
}

