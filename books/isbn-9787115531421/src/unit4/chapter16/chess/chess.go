package main

import "fmt"


func display(board [8][8]rune) {
	for _, row := range board {
		for _, column := range row {
			if column == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c", column)
			}
		}
		fmt.Println()
	}
}

func initialize() [8][8]rune {
	var board [8][8]rune
	
	blackChesses := [...]rune {
		'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r',
	}
	whiteChesses := [...]rune {
		'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R',
	}
	for i:= 0; i < 8; i++ {
		// black
		board[0][i] = blackChesses[i]
		board[1][i] = 'p'

		// white
		board[7][i] = whiteChesses[i]
		board[6][i] = 'P'
	}

	return board
}

func main() {
	board := initialize()
	display(board)
}