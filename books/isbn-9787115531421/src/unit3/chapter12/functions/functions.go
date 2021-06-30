package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64 {
	c := kelvinToCelsius(k)
	return celsiusToFahrenheit(c)
}

func main() {
	kelvin := 233.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "°K = ", celsius, "°C")
	fmt.Println()
	fmt.Print("0°K = ", kelvinToFahrenheit(0), "°F")
}