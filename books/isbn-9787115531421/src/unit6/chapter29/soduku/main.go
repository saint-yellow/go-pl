package main

import (
	"fmt"
	"os"
)

func main() {
	// initialize a grid
	digits := [rows][columns]int8 {
		{5,3,0,0,7,0,0,0,0},
		{6,0,0,0,1,9,5,0,0},
		{0,9,8,0,0,0,0,6,0},
		{8,0,0,0,6,0,0,0,3},
		{4,0,0,8,0,3,0,0,1},
		{7,0,0,0,2,6,0,0,6},
		{0,6,0,0,0,0,0,8,0},
		{0,0,0,4,1,9,0,0,5},
		{0,0,0,0,8,0,0,7,9},
	}
	sudoku := newSudoku(digits)
	sudoku.display()

	// set a cell but do not fix it
	err := sudoku.set(4, 4, 9, false)
	printError(err)
	sudoku.display()

	// set and fix a cell
	err = sudoku.set(3, 3, 2, true)
	printError(err)
	sudoku.display()

	// clear unfixed cell
	err = sudoku.clear(4, 4)
	printError(err)
	sudoku.display()

	// clear fixed cell
	err = sudoku.clear(3, 3)
	printError(err)
	sudoku.display()
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}