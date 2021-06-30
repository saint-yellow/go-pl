package main

import "fmt"
func main() {
	fmt.Println("Malacandra is not far. We can get there within 28 days.")
	var days = 28 // 28 days
	const hoursOfDay = 24 // Everyday has 24 hours.
	var L = 56000000 // 56000000 km
	var v = L / (days * 24)
	fmt.Printf("So the velocity of the spacecraft should be %v km/hour.\n", v)
}