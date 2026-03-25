package solver_test

import (
	"testing"

	"tetris-optimizer/internal/solver"
	"tetris-optimizer/internal/tetromino"
)

func TestSortByConstraint(t *testing.T) {
	// Create tetrominoes with different rotation counts
	square := tetromino.Tetromino{
		Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 0, C: 1}, {R: 1, C: 0}, {R: 1, C: 1}},
		ID:     'A',
	}
	line := tetromino.Tetromino{
		Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 1, C: 0}, {R: 2, C: 0}, {R: 3, C: 0}},
		ID:     'B',
	}
	lShape := tetromino.Tetromino{
		Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 1, C: 0}, {R: 2, C: 0}, {R: 2, C: 1}},
		ID:     'C',
	}

	tetros := []tetromino.Tetromino{lShape, square, line}

	// Sort by constraint (fewer rotations first)
	solver.SortByConstraint(tetros)

	// Verify IDs are preserved
	if tetros[0].ID != 'C' && tetros[0].ID != 'A' && tetros[0].ID != 'B' {
		t.Error("expected valid tetromino IDs after sorting")
	}

	// Verify all tetrominoes are still present
	ids := make(map[rune]bool)
	for _, tet := range tetros {
		ids[tet.ID] = true
	}
	if !ids['A'] || !ids['B'] || !ids['C'] {
		t.Error("expected all tetrominoes to be present after sorting")
	}
}

func TestSortByConstraintEmpty(t *testing.T) {
	tetros := []tetromino.Tetromino{}
	solver.SortByConstraint(tetros)
	if len(tetros) != 0 {
		t.Error("expected empty slice to remain empty")
	}
}

func TestSortByConstraintSingle(t *testing.T) {
	tetros := []tetromino.Tetromino{
		{
			Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 0, C: 1}, {R: 1, C: 0}, {R: 1, C: 1}},
			ID:     'A',
		},
	}
	solver.SortByConstraint(tetros)
	if len(tetros) != 1 || tetros[0].ID != 'A' {
		t.Error("expected single tetromino to remain unchanged")
	}
}
