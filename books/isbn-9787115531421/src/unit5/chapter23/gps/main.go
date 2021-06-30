package main

import (
	"fmt"
	"isbn-9787115531421/src/unit5/chapter22/distance/world"
	"isbn-9787115531421/src/unit5/chapter22/landing/location"
)

type gps struct {
	world world.World
	current location.Location
	destinatin location.Location
}

func (g gps) distance() float64 {
	return g.world.Distance(g.current, g.destinatin)
}

func (g gps) message() string {
	return fmt.Sprintf("%.2f km from %v to %v", g.distance(), g.current.String(), g.destinatin.String())
}

type rover struct {
	gps
}


func main() {
	mars := world.New(3389.5)
	bradbury := location.Location {
		Name:"Bradbury Landing", 
		Latitude: -4.5895, 
		Longitude:137.4417,
	}
	elysium := location.Location{
		Name:"Elysium Planitia", 
		Latitude:4.5, 
		Longitude:135.9,
	}

	gps := gps {
		world: mars,
		current: bradbury,
		destinatin: elysium,
	}

	curiosity := rover {
		gps: gps,
	}

	// method's forwarding
	fmt.Println(curiosity.message())
}