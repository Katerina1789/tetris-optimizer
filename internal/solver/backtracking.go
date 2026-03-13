package solver

import (
	"tetris-optimizer/internal/board"
	"tetris-optimizer/internal/tetromino"
)

// Backtrack attempts to place tetros[idx:] on the board and returns true when all tetrominoes are placed.
func Backtrack(b *board.Board, rotSets [][]tetromino.Tetromino, idx int) bool {
	if idx == len(rotSets) {
		return true
	}

	// Find first empty cell
	var fr, fc int
	found := false
	for r := 0; r < b.Size && !found; r++ {
		for c := 0; c < b.Size && !found; c++ {
			if b.Cells[r][c] == '.' {
				fr, fc = r, c
				found = true
			}
		}
	}

	if !found {
		return false
	}

	rots := rotSets[idx]

	// Try placing the tetromino with its origin near (fr, fc)
	for _, rt := range rots {
		for dr := -3; dr <= 3; dr++ {
			for dc := -3; dc <= 3; dc++ {
				baseR := fr + dr
				baseC := fc + dc

				if b.Place(rt, baseR, baseC) {
					if Backtrack(b, rotSets, idx+1) {
						return true
					}
					b.Remove(rt, baseR, baseC)
				}
			}
		}
	}

	return false
}

// QuickPrune checks if continuing is pointless.
func QuickPrune(b *board.Board, remaining int) bool {
	empty := 0
	for r := 0; r < b.Size; r++ {
		for c := 0; c < b.Size; c++ {
			if b.Cells[r][c] == '.' {
				empty++
			}
		}
	}

	// Only safe prune: not enough space left for remaining tetrominoes
	if empty < remaining*4 {
		return true
	}

	return false
}
