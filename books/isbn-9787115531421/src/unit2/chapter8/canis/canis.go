package main

import (
	"fmt"
)

func main() {
	const ( 
		lightSpeed = 299792
		secondsPerDay = 86400
		daysPerYear = 365
	)
	const distanceInKM = 236000000000000000
	const distanceInLY = distanceInKM / lightSpeed / secondsPerDay / daysPerYear

	fmt.Println("The Canis Major Dwarf is", distanceInKM, "KM away from the Earth.")
	fmt.Println("In other words, it is", distanceInLY, "LY.")
}