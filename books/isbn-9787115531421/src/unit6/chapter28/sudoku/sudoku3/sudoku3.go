package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	var g Grid
	err := g.Set(10, 0, 10)
	if err != nil {
		if errors, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errors))
			for _, e := range errors {
				fmt.Printf("- %v\n", e)
			}
		}
		os.Exit(1)
	}
}

type SudokuError []error

var (
	ErrBounds, ErrDigit = errors.New("out of bounds"), errors.New("invalid digit")
)

const rows, columns = 9, 9

type Grid [rows][columns]int8

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func (g *Grid) Set(row, column int, digit int8) error {
	var errors SudokuError
	
	if !inBounds(row, column) {
		errors = append(errors, ErrBounds)
	}

	if !validDigit(digit) {
		errors = append(errors, ErrDigit)
	}

	if len(errors) > 0 {
		return errors
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