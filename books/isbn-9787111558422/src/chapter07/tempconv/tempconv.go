package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", float64(c))
}

func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius((f-32)*5 / 9)
}

func KelvinToCelsius(k Kelvin) Celsius {
	return Celsius(k-273.15)
}

type celsiusFlag struct { Celsius }

func (f *celsiusFlag) Set(s string) error {
	var (
		unit string
		value float64
	)

	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FahrenheitToCelsius(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = KelvinToCelsius(Kelvin(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)

}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}