package main

import (
	"fmt"
)


type celsius float64
type fahrenheit float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

const (
	line = "======================="
	rowFormat = "| %8s | %8s |\n"
	numberFormat = "%.1f"
)


type getRowFn func(row int) (string, string)

func drawTable(header1 string, header2 string, rows int, getRow getRowFn) {

	// table header
	fmt.Println(line)
	fmt.Printf(rowFormat, header1, header2)
	fmt.Println(line)


	// table content (body)
	for row := 0; row < rows; row++ {
		cell1, cell2 := getRow(row)
		fmt.Printf(rowFormat, cell1, cell2)
	}


	// table footer
	fmt.Println(line)
}

func celsiusToFahrenheit(row int) (string, string) {
	c := celsius(row * 5 - 40)
	f := c.fahrenheit()
	cell1 := fmt.Sprintf(numberFormat, c)
	cell2 := fmt.Sprintf(numberFormat, f)
	return cell1, cell2
}

func fahrenheitToCelsius(row int) (string, string) {
	f := fahrenheit(row * 5 - 40)
	c := f.celsius()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}

func main() {
	drawTable("°C", "°F", 29, celsiusToFahrenheit)
	fmt.Println()
	drawTable("°F", "°C", 29, fahrenheitToCelsius)
}