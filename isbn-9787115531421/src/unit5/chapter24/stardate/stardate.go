package main

import (
	"fmt"
	"time"
)

func stardate(t time.Time) float64 {
	day := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + day + h
}

func main() {
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))
}