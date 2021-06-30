package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

func measureTemperature(samples int, s sensor) {
	for i := 0; i < samples; i++ {
		k := s()
		fmt.Printf("%v °K\n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151)+150)
}

func realSensor() kelvin {
	return 0
}

func main() {
	measureTemperature(3, fakeSensor)
	measureTemperature(3, realSensor)
}