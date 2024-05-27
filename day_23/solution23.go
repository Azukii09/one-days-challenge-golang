package day_23

import (
	"strconv"
	"strings"
)

const mineChar = '*'
const emptyChar = ' '

// Annotate returns an annotated board
func Annotate(board []string) []string {
	output := make([]string, len(board))
	for row := range board {
		output[row] = fillRow(board, row)
	}
	return output
}
func fillRow(board []string, row int) string {
	var result strings.Builder
	for column, character := range board[row] {
		count := 0
		if character == emptyChar {
			count = countAdjacentMines(board, row, column)
		}
		if count > 0 {
			result.WriteString(strconv.Itoa(count))
		} else {
			result.WriteRune(character)
		}
	}
	return result.String()
}
func countAdjacentMines(board []string, row int, column int) int {
	count := 0
	minRow := max(row-1, 0)
	maxRow := min(row+1, len(board)-1)
	minCol := max(column-1, 0)
	maxCol := min(column+1, len(board[row])-1)
	for _, s := range board[minRow : maxRow+1] {
		for _, r := range s[minCol : maxCol+1] {
			if r == mineChar {
				count++
			}
		}
	}
	return count
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
