package main

import "fmt"

func main() {
	var k kelvin = 234.0
	c := k.celsius()
	f := c.fahrenheit()

	fmt.Println(k, "°K = ", c, "°C = ", f, "°F")
	fmt.Println(f.kelvin(), "°K = ", k.celsius(), "°C = ", c.fahrenheit(), "°F")
	fmt.Println(c.kelvin(), "°K = ", f.celsius(), "°C = ", k.fahrenheit(), "°F")
}