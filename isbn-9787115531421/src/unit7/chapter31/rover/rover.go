package main

import (
	"image"
	"log"
	"time"
)

type roverDriver struct {
	cmd chan command
}

type command int

const (
	right = command(0)
	left = command(1)
	start = command(2)
	stop = command(3)
)

func newRoverDriver() *roverDriver {
	r := &roverDriver{cmd: make(chan command)}
	go r.drive()
	return r
}

func (r *roverDriver) drive() {
	pos := image.Point{X:0, Y:0}
	direction := image.Point{X:1, Y:0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	speed := 1
	for {
		select {
			case c := <-r.cmd:
				switch c {
					case right:
						direction = image.Point{X:-direction.X, Y:direction.X}
					case left:
						direction = image.Point{X:direction.Y, Y:-direction.X}
					case start:
						speed = 1
					case stop:
						speed = 0
				}
				log.Printf("new direction %v; speed: %d", direction, speed)
			case <-nextMove:
				pos = pos.Add(direction.Mul(speed))
				log.Printf("moved to %v", pos)
				nextMove = time.After(updateInterval)
		}
	}
}

func (r *roverDriver) left() {
	r.cmd <- left
}

func (r *roverDriver) right() {
	r.cmd <- right
}

func (r *roverDriver) start() {
	r.cmd <- start
}

func (r * roverDriver) stop() {
	r.cmd <- stop
}