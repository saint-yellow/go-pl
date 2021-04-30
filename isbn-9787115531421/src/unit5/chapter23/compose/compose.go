package main

import "fmt"

type report struct {
	sol int
	temperature temperature
	location location
}

func (r report) average() celsius {
	return r.temperature.average()
}

type temperature struct {
	high, low celsius
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

type location struct {
	latitude, longitude float64 
}

type celsius float64

func main() {
	bradbury := location{-4.5895, 137.4417}
	temperature := temperature{high: -1.0, low: -78.0}
	fmt.Printf("average %v °C\n", temperature.average())

	report := report{sol: 15, temperature: temperature, location: bradbury}

	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v °C\n", report.temperature.high)
}