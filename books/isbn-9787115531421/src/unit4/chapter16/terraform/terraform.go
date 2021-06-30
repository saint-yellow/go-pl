package main

import "fmt"

func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "New-" + planets[i]
	}
	fmt.Println(planets)
}

func main() {
	planets := [...]string {
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	terraform(planets)
	fmt.Println(planets)
}