package main

import (
	"image"
	"math/rand"
	"sync"
)

type marsGrid struct {
	bounds image.Rectangle
	mu sync.Mutex
	cells [][]cell

}

func newMarsGrid(size image.Point) *marsGrid {
	grid := &marsGrid{
		bounds: image.Rectangle{Max:size},
		cells: make([][]cell, size.Y),
	}
	for y := range grid.cells {
		grid.cells[y] = make([]cell, size.X)
		for x := range grid.cells[y] {
			cell := &grid.cells[y][x]
			cell.groupData.lifeSigns = rand.Intn(1000)
		}
	}
	return grid
}

func (g *marsGrid) size() image.Point {
	return g.bounds.Max
}

func (g *marsGrid) occupy(p image.Point) *occupier {
	g.mu.Lock()
	defer g.mu.Unlock()

	cell := g.cell(p)
	if cell == nil || cell.ocp != nil {
		return nil
	}
	cell.ocp = &occupier{grid: g, pos: p}
	return cell.ocp
}

func (g *marsGrid) cell(p image.Point) *cell {
	if !p.In(g.bounds) {
		return nil
	}
	return &g.cells[p.X][p.Y]
}