package main

import (
	"image"
	"log"
	"math/rand"
	"time"
)

type roverDriver struct {
	cmd chan command
	ocp *occupier
	name string
	rdo *radio
}

type command int

const (
	right command = 0
	left command = 1
)

func newRoverDriver(name string, ocp *occupier, marsToEarth chan []message) *roverDriver {
	r := &roverDriver{
		cmd: make(chan command),
		ocp: ocp,
		name: name,
		rdo: newRadio(marsToEarth),
	}
	go r.drive()
	return r
}

func (r *roverDriver) drive() {
	log.Printf("%s initial position %v", r.name, r.ocp.pos)
	direction := image.Point{X:1, Y:0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
			case c := <-r.cmd:
				switch c {
					case right:
						direction = image.Point{X:-direction.X, Y:direction.X}
					case left:
						direction = image.Point{X:direction.Y, Y:-direction.X}
				}
				log.Printf("new direction %v; direction: %d", direction, direction)
			case <-nextMove:
				nextMove = time.After(updateInterval)
				newPos := r.ocp.pos.Add(direction)
				if r.ocp.moveTo(newPos) {
					log.Printf("%s moved to %v", r.name, newPos)
					r.checkForLife()
					break
				}
				log.Printf("%s is blocked trying to move from %v to %v", r.name, r.ocp.pos, newPos)
				randomDirection := rand.Intn(3) + 1
				for i := 0; i < randomDirection; i++ {
					direction = image.Point{X: -direction.Y, Y: direction.X}
				}
				log.Printf("%s new random direction %v", r.name, direction)
		}
	}
}

func (r *roverDriver) checkForLife() {
	sensorData := r.ocp.sense()
	if sensorData.lifeSigns < 900 {
		return
	}
	r.rdo.sendToEarth(message{
		pos: r.ocp.pos,
		lifeSigns: sensorData.lifeSigns,
		rover: r.name,
	})
}

// func (r *roverDriver) left() {
// 	r.cmd <- left
// }

// func (r *roverDriver) right() {
// 	r.cmd <- right
// }

func startRoverDriver(name string, grid *marsGrid, marsToEarth chan []message) *roverDriver {
	var o *occupier
	for o == nil {
		startPoint := image.Point{X: rand.Intn(grid.size().X), Y: rand.Intn(grid.size().Y)}
		o = grid.occupy(startPoint)
	}
	return newRoverDriver(name, o, marsToEarth)
}