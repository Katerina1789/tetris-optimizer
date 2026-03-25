package board

import "tetris-optimizer/internal/tetromino"

// Place tries to place a tetromino on the board at the given base position.
// Returns true if placement succeeded, false otherwise.
func (board *Board) Place(tetrominoToPlace tetromino.Tetromino, baseRow, baseColumn int) bool {
	// First pass: check if all blocks can be placed
	for _, point := range tetrominoToPlace.Blocks {
		row := baseRow + point.Row
		column := baseColumn + point.Column
		if !board.InBounds(row, column) || !board.IsEmpty(row, column) {
			return false
		}
	}

	// Second pass: place all blocks (we know it's safe now)
	for _, point := range tetrominoToPlace.Blocks {
		row := baseRow + point.Row
		column := baseColumn + point.Column
		board.setCell(row, column, tetrominoToPlace.ID)
	}

	return true
}

// Remove clears a tetromino from the board at the given base position
func (board *Board) Remove(tetrominoToRemove tetromino.Tetromino, baseRow, baseColumn int) {
	for _, point := range tetrominoToRemove.Blocks {
		row := baseRow + point.Row
		column := baseColumn + point.Column
		if board.InBounds(row, column) && board.Cell(row, column) == tetrominoToRemove.ID {
			board.setCell(row, column, '.')
		}
	}
}

// FirstEmpty finds the first empty cell on the board (scanning left-to-right, top-to-bottom).
// Returns (row, column, true) if found, or (0, 0, false) if board is full.
func (board *Board) FirstEmpty() (int, int, bool) {
	for row := 0; row < board.Size; row++ {
		for column := 0; column < board.Size; column++ {
			if board.Cell(row, column) == '.' {
				return row, column, true
			}
		}
	}
	return 0, 0, false
}
