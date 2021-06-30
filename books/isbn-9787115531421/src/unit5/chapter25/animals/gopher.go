package main

import "math/rand"

type gopher struct {
	name string
}

func (g gopher) String() string {
	return g.name
}

func (g gopher) move() string {
	switch rand.Intn(2) {
		case 0:
			return "scurries along the ground"
		default:
			return "burrows in the sand"
	}
}

func (g gopher) eat() string {
	switch rand.Intn(5) {
		case 0:
			return "carrot"
		case 1:
			return "lettuce"
		case 2:
			return "radish"
		case 3:
			return "corn"
	    default:
			return "root"
	}
}