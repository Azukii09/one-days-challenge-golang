package day_26

import "errors"

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

// StateOfTicTacToe determines the state of the Tic-Tac-Toe game.
func StateOfTicTacToe(board []string) (State, error) {
	if len(board) != 3 {
		return "", errors.New("invalid board size")
	}
	for _, row := range board {
		if len(row) != 3 {
			return "", errors.New("invalid board size")
		}
	}
	// Convert board to a flat representation
	flatBoard := make([]string, 9)
	xCount, oCount := 0, 0
	for i, row := range board {
		for j, cell := range row {
			flatBoard[i*3+j] = string(cell)
			if cell == 'X' {
				xCount++
			} else if cell == 'O' {
				oCount++
			}
		}
	}
	if oCount > xCount || xCount > oCount+1 {
		return "", errors.New("invalid number of moves")
	}
	lines := [][]int{
		// Horizontal lines
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		// Vertical lines
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		// Diagonal lines
		{0, 4, 8}, {2, 4, 6},
	}
	xWins, oWins := false, false
	// Check for a win
	for _, line := range lines {
		if flatBoard[line[0]] != " " && flatBoard[line[0]] == flatBoard[line[1]] && flatBoard[line[1]] == flatBoard[line[2]] {
			if flatBoard[line[0]] == "X" {
				xWins = true
			} else if flatBoard[line[0]] == "O" {
				oWins = true
			}
		}
	}
	if xWins && oWins {
		return "", errors.New("invalid board state: both players cannot win")
	}
	if xWins || oWins {
		if (xWins && xCount == oCount+1) || (oWins && xCount == oCount) {
			return Win, nil
		}
		return "", errors.New("invalid board state: incorrect number of moves for a win")
	}
	// Check for ongoing or draw
	for _, cell := range flatBoard {
		if cell == " " {
			return Ongoing, nil
		}
	}
	return Draw, nil
}
