package main

import "math/rand"


type honeyBee struct {
	name string
}

func (hb honeyBee) String() string {
	return hb.name
}

func (hb honeyBee) move() string {
	switch rand.Intn(2) {
		case 0:
			return "buzzes about"
		default:
			return "files to infinity and beyond"
	}
}

func (hb honeyBee) eat() string {
	switch rand.Intn(2) {
		case 0:
			return "pollen"
		default:
			return "nectar"
	}
}

