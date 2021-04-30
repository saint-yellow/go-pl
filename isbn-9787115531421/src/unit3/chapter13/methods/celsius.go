package main

type celsius float64

func (c celsius) kelvin() kelvin {
	return kelvin(c+273.15)
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}