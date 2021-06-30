package main

import "fmt"

type report struct {
	sol
	temperature
	location
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

type temperature struct {
	high, low celsius
}

type location struct {
	latitude, longitude float64 
}

type sol int

type celsius float64

func main() {
	report := report{sol: 15}

	fmt.Println(report.days(1446))
	fmt.Println(report.sol.days(1446))
}