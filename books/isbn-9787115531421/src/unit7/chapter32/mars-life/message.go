package main

import (
	"image"
	"log"
	"time"
)

type message struct {
	pos image.Point
	lifeSigns int
	rover string
}

func earthReceiver(messageChannel chan []message) {
	for {
		time.Sleep(dayLength - receiveTimePerDay)
		receiveMarsMessage(messageChannel)
	}
}

func receiveMarsMessage(messageChannel chan []message) {
	finished := time.After(receiveTimePerDay)
	for {
		select {
		case <-finished:
			return
		case ms := <-messageChannel:
			for _, m := range ms {
				log.Printf("earth received report of life sign level %d from %s at %v", m.lifeSigns, m.rover, m.pos)
			}
		}
	}
}

