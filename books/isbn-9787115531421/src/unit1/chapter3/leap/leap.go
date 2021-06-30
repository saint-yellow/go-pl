package main

import (
	"fmt"
)
func main() {
	var year = 2100
	fmt.Printf("The year is %v, should you leap?\n", year)
	var leap = isLeap(year)
	if leap {
		fmt.Println("Look before you leap!")
	} else {
		fmt.Println("Keep your feet on the ground.")
	}
}


func isLeap(year int) bool {
	var leap = year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)
	return leap
}