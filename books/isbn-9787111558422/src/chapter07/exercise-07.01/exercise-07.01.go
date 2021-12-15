package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	s.Split(bufio.ScanWords)
	for s.Scan() {
		*c++
	}
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		*c++
	}
	return len(p), nil
}

func main() {
	s := "Saint Yellow\nSaint-Yellow"

	var c1 WordCounter
	c1.Write([]byte(s))
	fmt.Println(c1)

	var c2 LineCounter
	c2.Write([]byte(s))
	fmt.Println(c2)

	// var c3 Counter
	// c3.Write([]byte(s))
	// fmt.Println(c3)

	// c3 = 0
	// SplitType = 1
	// c3.Write([]byte(s))
	// fmt.Println(c3)
}

// type Counter int
// var SplitType int

// func (c *Counter) Write(p []byte) (int, error) {
// 	s := bufio.NewScanner(bytes.NewReader(p))
// 	switch SplitType {
// 	case 0:
// 		s.Split(bufio.ScanWords)
// 	default:
// 		s.Split(bufio.ScanLines)
// 	}
// 	for s.Scan() {
// 		*c++
// 	}
// 	return len(p), nil
// }