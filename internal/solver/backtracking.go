package solver

import (
	"tetris-optimizer/internal/board"
	"tetris-optimizer/internal/tetromino"
)

const (
	// MaxTetrominoSize is the maximum dimension of a tetromino (4x4 grid)
	// We search from -3 to +3 relative to the first empty cell to ensure
	// we can place any tetromino block at that position
	MaxTetrominoSize = 3
)

// Backtrack attempts to place all tetrominoes on the board using recursive backtracking.
// It returns true when all tetrominoes are successfully placed.
func Backtrack(currentBoard *board.Board, rotationSets [][]tetromino.Tetromino, tetrominoIndex int) bool {
	// Base case: all tetrominoes placed successfully
	if tetrominoIndex == len(rotationSets) {
		return true
	}

	// Find first empty cell on board (optimization: ensures no gaps)
	firstEmptyRow, firstEmptyColumn, found := currentBoard.FirstEmpty()
	if !found {
		// Board is full but not all tetrominoes placed
		return false
	}

	// Get all rotations for current tetromino
	rotations := rotationSets[tetrominoIndex]

	// Try each rotation of the current tetromino
	for _, rotation := range rotations {
		// Try placing tetromino at different positions near the first empty cell
		// We need to check positions from -3 to +3 because a tetromino can be up to 4x4,
		// and we want to try placing it so that any of its blocks covers the empty cell
		for deltaRow := -MaxTetrominoSize; deltaRow <= MaxTetrominoSize; deltaRow++ {
			for deltaColumn := -MaxTetrominoSize; deltaColumn <= MaxTetrominoSize; deltaColumn++ {
				baseRow := firstEmptyRow + deltaRow
				baseColumn := firstEmptyColumn + deltaColumn

				// Try to place tetromino at this position
				if currentBoard.Place(rotation, baseRow, baseColumn) {
					// Recursively try to place remaining tetrominoes
					if Backtrack(currentBoard, rotationSets, tetrominoIndex+1) {
						return true // Solution found!
					}
					// Backtrack: remove tetromino and try next position
					currentBoard.Remove(rotation, baseRow, baseColumn)
				}
			}
		}
	}

	// No valid placement found for this tetromino
	return false
}
