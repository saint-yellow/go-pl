package world

import (
	"isbn-9787115531421/src/unit5/chapter22/landing/location"
	"math"
)

type World struct {
	Radius float64
}

func New(radius float64) World {
	return World{radius}
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w World) Distance(p1, p2 location.Location) float64 {
	s1, c1 := math.Sincos(rad(p1.Latitude))
	s2, c2 := math.Sincos(rad(p2.Latitude))
	clong := math.Cos(rad(p1.Longitude - p2.Longitude))
	return w.Radius * math.Acos(s1*s2+c1*c2*clong)
}