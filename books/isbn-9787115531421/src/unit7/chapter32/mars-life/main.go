package main

import (
	"fmt"
	"image"
	"time"
)

const (
	dayLength = 24 * time.Second
	receiveTimePerDay = 2 * time.Second
)

func main() {
	marsToEarth := make(chan []message)
	go earthReceiver(marsToEarth)

	gridSize := image.Point{X:20, Y:20}
	grid := newMarsGrid(gridSize)

	rover := make([]*roverDriver, 5)
	for i := range rover {
		rover[i] = startRoverDriver(fmt.Sprint("rover", i), grid, marsToEarth)
	}
	time.Sleep(60 * time.Second)
}