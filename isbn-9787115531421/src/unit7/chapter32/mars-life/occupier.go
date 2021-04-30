package main

import "image"

type occupier struct {
	grid *marsGrid
	pos image.Point
}

func (o *occupier) moveTo(p image.Point) bool {
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()

	newCell := o.grid.cell(p)
	if newCell == nil || newCell.ocp != nil {
		return false
	}
	o.grid.cell(o.pos).ocp = nil
	newCell.ocp = o
	o.pos = p
	return true
}

func (o *occupier) sense() sensorData {
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()

	return o.grid.cell(o.pos).groupData
}
