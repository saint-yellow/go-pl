package main

import "fmt"

func main() {
	var bh float64 = 32767
	bh++ // wrap around: bh > math.MaxInt16
	var h = int16(bh)

	fmt.Println(h)
}