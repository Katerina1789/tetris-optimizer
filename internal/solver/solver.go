package solver

import (
	"fmt"
	"math"

	"tetris-optimizer/internal/board"
	"tetris-optimizer/internal/tetromino"
)

// Solve tries to place all tetrominoes into the smallest square boardand increases the board size until a solution is found or a limit is reached.
func Solve(tetros []tetromino.Tetromino) (*board.Board, error) {
	n := len(tetros)
	minSize := int(math.Ceil(math.Sqrt(float64(n * 4))))

	// Sort tetrominoes by constraint (fewer rotations first)
	// sortByConstraint(tetros)  // ← comment this out for now

	// Precompute rotations once
	rotSets := make([][]tetromino.Tetromino, len(tetros))
	for i, t := range tetros {
		rotSets[i] = tetromino.Rotations(t)
	}

	for size := minSize; size <= 20; size++ {
		b := board.New(size)
		if Backtrack(b, rotSets, 0) {
			return b, nil
		}
	}
	return nil, fmt.Errorf("no solution")
}
