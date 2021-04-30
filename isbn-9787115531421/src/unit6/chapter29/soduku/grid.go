package main

import (
	"errors"
	"fmt"
)

const (
	rows, columns = 9, 9
	empty = 0
)

var (
	errBounds = errors.New("out of bounds")
	errDigit = errors.New("invalid digit")
	errInRow = errors.New("digit already present in the row")
	errInColumn = errors.New("digit already present in this column")
	errInRegion = errors.New("digit already in the region")
	errFixedDigit = errors.New("invalid digits cannot be overwritten")
)

type cell struct {
	digit int8
	fixed bool
}

type grid [rows][columns]cell

func newSudoku(digits [rows][columns]int8) *grid {
	var g grid
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			d := digits[r][c]
			if d != empty {
				g[r][c].digit = d
				g[r][c].fixed = true
			}
		}
	}
	return &g
}

func (g *grid) set(row, column int, digit int8, fixed bool) error {
	switch {
		case !inBounds(row, column):
			return errBounds
		case !validDigit(digit):
			return errDigit
		case g.isFixed(row, column):
			return errFixedDigit
		case g.inRow(row, digit):
			return errInRow
		case g.inColumn(column, digit):
			return errInColumn
		case g.inRegion(row, column, digit):
			return errInRegion
	}

	g[row][column].digit = digit
	g[row][column].fixed = fixed
	return nil
}

func (g *grid) clear(row, column int) error {
	switch {
	case !inBounds(row, column):
		return errBounds
	case g.isFixed(row, column):
		return errFixedDigit
	}

	g[row][column].digit = empty
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows || column < 0 || column > columns {
		return false
	}

	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func (g *grid) inRow(row int, digit int8) bool {
	for c := 0; c < columns; c++ {
		if g[row][c].digit == digit {
			return true
		}
	}
	return false
}

func (g *grid) inColumn(column int, digit int8) bool {
	for r := 0; r < rows; r++ {
		if g[r][column].digit == digit {
			return true
		}
	}
	return false
}

func (g *grid) inRegion(row, column int, digit int8) bool {
	startRow, startColumn := row/3*3, column/3*3
	for r := startRow; r < startRow+3; r++ {
		for c := startColumn; c < startColumn; c++ {
			if g[r][c].digit == digit {
				return true
			}
		}
	}
	return false
}

func (g *grid) isFixed(row, column int) bool {
	return g[row][column].fixed
}

func (g grid) display() {
	redColor := "\033[31m"
	blueColor := "\033[34m"
	whiteColor := "\033[37m"

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			cell := g[row][column]
			var color string
			if cell.fixed {
				color = string(redColor)
			} else {
				color = string(blueColor)
			}

			if column == 8 {
				fmt.Print(color, " ", cell.digit, " ", "\n")
			} else {
				fmt.Print(color, " ", cell.digit, " ")
			}
		}
	}

	fmt.Println(string(whiteColor))
}