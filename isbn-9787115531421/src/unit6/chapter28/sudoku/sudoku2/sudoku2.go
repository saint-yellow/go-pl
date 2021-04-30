package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit = errors.New("invalid digit")
)

const rows, columns = 9, 9

type Grid [rows][columns]int8

func main() {
	var g Grid
	err := g.Set(10, 0, 5)
	if err != nil {
		switch err {
			case ErrBounds, ErrDigit:
				fmt.Println("Out of bounds or invalid digit")
			default:
				fmt.Println(err)
		}
		os.Exit(1)
	}
}

func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return ErrBounds
	}

	if !validDigit(digit) {
		return ErrDigit
	}

	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}

	if column < 0 || column >= columns {
		return false
	}

	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}