package main

import (
	"fmt"
	"strings"
)

var t interface {
	talk() string
}

type martian struct {}

func (m martian) talk() string {
	return "neck neck"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func main() {
	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())
}