package main

import "fmt"

type turtle struct {
	x, y int
}

func (t *turtle) up() {
	t.y++
}

func (t *turtle) down() {
	t.y--
}

func (t *turtle) left() {
	t.x--
}

func (t *turtle) right() {
	t.x++
}

func (t turtle) String() string {
	return fmt.Sprintf("(%d, %d)", t.x, t.y)
}