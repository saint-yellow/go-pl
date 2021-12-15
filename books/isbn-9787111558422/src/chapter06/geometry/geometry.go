package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p, q := Point{1, 2}, Point{4, 6}
	fmt.Println(Distance(p,q))
	fmt.Println(p.Distance(q))

	path := Path{p, q, {3, 5}}
	fmt.Println(path.Distance())

	p.ScaleBy(2)
	fmt.Println(p)

	s := &Point{2, 3}
	fmt.Println(s)
	s.ScaleBy(3)
	fmt.Println(*s)
}