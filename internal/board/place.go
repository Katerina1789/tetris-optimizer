package board

// Placement helpers: place and remove tetrominoes on the board.

import "tetris-optimizer/internal/tetromino"

// Place tries to place tetromino t with its origin at (baseR, baseC).
// Returns true if placement succeeded.
func (b *Board) Place(t tetromino.Tetromino, baseR, baseC int) bool {
	for _, p := range t.Blocks {
		r := baseR + p.R
		c := baseC + p.C
		if !b.InBounds(r, c) || !b.IsEmpty(r, c) {
			return false
		}
	}
	for _, p := range t.Blocks {
		r := baseR + p.R
		c := baseC + p.C
		b.Cells[r][c] = t.ID
	}
	return true
}

// Remove clears tetromino t from the board at origin (baseR, baseC).
func (b *Board) Remove(t tetromino.Tetromino, baseR, baseC int) {
	for _, p := range t.Blocks {
		r := baseR + p.R
		c := baseC + p.C
		if b.InBounds(r, c) && b.Cells[r][c] == t.ID {
			b.Cells[r][c] = '.'
		}
	}
}
