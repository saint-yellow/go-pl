package universe

import (
	"fmt"
	"math/rand"
)


const (
	width = 80
	height = 15
)

type Universe [][]bool


func New() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

func (u Universe) set(x int, y int, b bool) {
	u[y][x] = b
}

func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		u.set(rand.Intn(width), rand.Intn(height), true)
	} 
}

func (u Universe) alive(x int, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

func (u Universe) neighbors(x int, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v ==0 && h == 0) && u.alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}

func (u Universe) next(x int, y int) bool {
	n := u.neighbors(x, y)
	return n == 3 || n == 2 && u.alive(x, y)
}

func (u Universe) string() string {
	var b byte
	buffer := make([]byte, 0, (width+1)*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b = ' '
			if u[y][x] {
				b = '*'
			}
			buffer = append(buffer, b)
		}
	}
	return string(buffer)
}

func (u Universe) Show() {
	fmt.Print("\x0c", u.string())
}

func Step(u1 Universe, u2 Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			u2.set(x, y, u1.next(x, y))
		}
	}
}

