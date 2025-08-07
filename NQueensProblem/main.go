package main

import (
	"fmt"
	"math"
)

func main() {
	backtrack(0, []space{}, 5)
}

func backtrack(row int, queens []space, n int) bool {
	if row == n {
		printQueens(queens)
	}

	for col := range n {
		if isValid(row, col, queens) {
			queens = append(queens, space{row: row, col: col})
			if backtrack(row+1, queens, n) {
				return true
			}
			queens = queens[:len(queens)-1]
		}
	}

	return true
}

func isValid(row int, col int, queens []space) bool {
	for _, queen := range queens {
		if queen.col == col || math.Abs(float64(queen.col-col)) == math.Abs(float64(queen.row-row)) {
			return false
		}
	}
	return true
}

func printQueens(queens []space) {
	fmt.Printf("Printing solution \n\n")
	for row := range len(queens) {
		for col := range len(queens) {
			if queens[row].col == col {
				fmt.Printf(" Q ")
			} else {
				fmt.Printf(" o ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n~~~~~~~~~~~~~~~~~~~~~\n")
}

type space struct {
	row int
	col int
}
