package solver

// Package solver implements the backtracking algorithm to fit all tetrominoes
// into the smallest possible square board.

import (
	"fmt"
	"math"

	"tetris-optimizer/internal/board"
	"tetris-optimizer/internal/tetromino"
)

// Solve tries to place all tetrominoes into the smallest square board.
// It increases the board size until a solution is found or a limit is reached.
func Solve(tetros []tetromino.Tetromino) (*board.Board, error) {
	n := len(tetros)
	minSize := int(math.Ceil(math.Sqrt(float64(n * 4))))

	for size := minSize; size <= 20; size++ { // 20 is a safe upper bound for typical inputs
		b := board.New(size)
		if backtrack(b, tetros, 0) {
			return b, nil
		}
	}
	return nil, fmt.Errorf("no solution")
}
