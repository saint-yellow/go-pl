package main

type kelvin float64

func (k kelvin) celsius() celsius {
	return celsius(k-273.15)
}

func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}