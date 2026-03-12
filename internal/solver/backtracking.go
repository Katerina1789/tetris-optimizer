package solver

// Backtracking search: tries to place tetrominoes one by one.

import (
	"tetris-optimizer/internal/board"
	"tetris-optimizer/internal/tetromino"
)

// backtrack attempts to place tetros[idx:] on the board.
// Returns true when all tetrominoes are placed.
func backtrack(b *board.Board, tetros []tetromino.Tetromino, idx int) bool {
	if idx == len(tetros) {
		return true
	}

	t := tetros[idx]
	rots := tetromino.Rotations(t)

	for r := 0; r < b.Size; r++ {
		for c := 0; c < b.Size; c++ {
			for _, rt := range rots {
				if b.Place(rt, r, c) {
					if backtrack(b, tetros, idx+1) {
						return true
					}
					b.Remove(rt, r, c)
				}
			}
		}
	}
	return false
}
