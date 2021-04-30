package main

import (
	"fmt"
	"image"
	"log"
	"time"
)

func printerWorker() {
	n := 0
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			n++
			fmt.Println(n)
			next = time.After(time.Second)
		}
	}
}

func positionerWorker() {
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X:1, Y:0}
	delay := time.Second
	next := time.After(delay)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			log.Println("current position is ", pos)
			delay += time.Second / 2
			next = time.After(delay)
		}
	}
}

func main() {
	go positionerWorker()
	time.Sleep(5 * time.Second)
}