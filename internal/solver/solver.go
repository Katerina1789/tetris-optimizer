package solver

import (
	"fmt"
	"math"

	"tetris-optimizer/internal/board"
	"tetris-optimizer/internal/tetromino"
)

const (
	// MaxTetrominoes is the maximum number of tetrominoes supported (A-Z)
	MaxTetrominoes = 26
	// MaxBoardSize is the maximum board size to try before giving up
	MaxBoardSize = 20
)

// Solve tries to place all tetrominoes into the smallest square board and increases the board size until a solution is found or a limit is reached.
func Solve(tetrominoes []tetromino.Tetromino) (*board.Board, error) {
	numberOfTetrominoes := len(tetrominoes)

	// Validate input
	if numberOfTetrominoes == 0 {
		return nil, fmt.Errorf("No tetrominoes provided")
	}
	if numberOfTetrominoes > MaxTetrominoes {
		return nil, fmt.Errorf("Too many tetrominoes: got %d, maximum is %d (A-Z)", numberOfTetrominoes, MaxTetrominoes)
	}

	// Apply heuristic: sort tetrominoes by constraint AND precompute rotations
	// This calculates rotations once and sorts by constraint (fewer rotations first)
	// Returns precomputed rotation sets for efficient backtracking
	rotationSets := SortByConstraint(tetrominoes)

	// Calculate minimum board size needed to fit all blocks
	minimumSize := int(math.Ceil(math.Sqrt(float64(numberOfTetrominoes * 4))))

	// Try increasing board sizes until solution found
	for size := minimumSize; size <= MaxBoardSize; size++ {
		currentBoard := board.New(size)
		if Backtrack(currentBoard, rotationSets, 0) {
			return currentBoard, nil
		}
	}

	return nil, fmt.Errorf("No solution found: unable to fit all tetrominoes on board up to size %dx%d", MaxBoardSize, MaxBoardSize)
}
