package solver_test

import (
	"testing"

	"tetris-optimizer/internal/solver"
	"tetris-optimizer/internal/tetromino"
)

func TestSortByConstraint(t *testing.T) {
	// Create tetrominoes with different rotation counts
	// Square: 1 unique rotation (most constrained)
	squareTetromino := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 0, Column: 1},
			{Row: 1, Column: 0},
			{Row: 1, Column: 1},
		},
		ID: 'A',
	}

	// Line: 2 unique rotations (horizontal and vertical)
	lineTetromino := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 1, Column: 0},
			{Row: 2, Column: 0},
			{Row: 3, Column: 0},
		},
		ID: 'B',
	}

	// L-shape: 4 unique rotations (least constrained)
	lShapeTetromino := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 1, Column: 0},
			{Row: 2, Column: 0},
			{Row: 2, Column: 1},
		},
		ID: 'C',
	}

	// Start with least constrained first (reverse order)
	tetrominoes := []tetromino.Tetromino{lShapeTetromino, lineTetromino, squareTetromino}

	// Sort by constraint (most constrained first)
	solver.SortByConstraint(tetrominoes)

	// Verify: Square (1 rotation) should be first
	// The exact order depends on rotation counts, but square should come before L-shape
	firstRotations := tetromino.Rotations(tetrominoes[0])
	lastRotations := tetromino.Rotations(tetrominoes[len(tetrominoes)-1])

	if len(firstRotations) > len(lastRotations) {
		t.Errorf("expected most constrained first, but first has %d rotations and last has %d",
			len(firstRotations), len(lastRotations))
	}

	// Verify all tetrominoes are still present
	idMap := make(map[rune]bool)
	for _, currentTetromino := range tetrominoes {
		idMap[currentTetromino.ID] = true
	}
	if !idMap['A'] || !idMap['B'] || !idMap['C'] {
		t.Error("expected all tetrominoes to be present after sorting")
	}
}

func TestSortByConstraintEmpty(t *testing.T) {
	tetrominoes := []tetromino.Tetromino{}
	solver.SortByConstraint(tetrominoes)
	if len(tetrominoes) != 0 {
		t.Error("expected empty slice to remain empty")
	}
}

func TestSortByConstraintSingle(t *testing.T) {
	tetrominoes := []tetromino.Tetromino{
		{
			Blocks: []tetromino.Point{
				{Row: 0, Column: 0},
				{Row: 0, Column: 1},
				{Row: 1, Column: 0},
				{Row: 1, Column: 1},
			},
			ID: 'A',
		},
	}

	solver.SortByConstraint(tetrominoes)

	if len(tetrominoes) != 1 || tetrominoes[0].ID != 'A' {
		t.Error("expected single tetromino to remain unchanged")
	}
}

func TestSortByConstraintStability(t *testing.T) {
	// Create multiple tetrominoes with same rotation count
	// Verify stable sort (relative order preserved)
	firstSquare := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 0, Column: 1},
			{Row: 1, Column: 0},
			{Row: 1, Column: 1},
		},
		ID: 'A',
	}

	secondSquare := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 0, Column: 1},
			{Row: 1, Column: 0},
			{Row: 1, Column: 1},
		},
		ID: 'B',
	}

	thirdSquare := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 0, Column: 1},
			{Row: 1, Column: 0},
			{Row: 1, Column: 1},
		},
		ID: 'C',
	}

	tetrominoes := []tetromino.Tetromino{firstSquare, secondSquare, thirdSquare}
	solver.SortByConstraint(tetrominoes)

	// All have same rotation count, so order should be preserved (stable sort)
	if tetrominoes[0].ID != 'A' || tetrominoes[1].ID != 'B' || tetrominoes[2].ID != 'C' {
		t.Error("expected stable sort to preserve relative order of equal elements")
	}
}

func TestSortByConstraintActualRotationCounts(t *testing.T) {
	// Test with actual rotation counts to verify sorting logic
	squareTetromino := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 0, Column: 1},
			{Row: 1, Column: 0},
			{Row: 1, Column: 1},
		},
		ID: 'S',
	}

	lShapeTetromino := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 1, Column: 0},
			{Row: 2, Column: 0},
			{Row: 2, Column: 1},
		},
		ID: 'L',
	}

	// Get actual rotation counts
	squareRotations := tetromino.Rotations(squareTetromino)
	lShapeRotations := tetromino.Rotations(lShapeTetromino)

	t.Logf("Square has %d unique rotations", len(squareRotations))
	t.Logf("L-shape has %d unique rotations", len(lShapeRotations))

	// Start with L-shape first (more rotations)
	tetrominoes := []tetromino.Tetromino{lShapeTetromino, squareTetromino}
	solver.SortByConstraint(tetrominoes)

	// After sorting, square (fewer rotations) should be first
	if len(squareRotations) < len(lShapeRotations) {
		if tetrominoes[0].ID != 'S' {
			t.Errorf("expected square (fewer rotations) to be first, but got %c", tetrominoes[0].ID)
		}
	}
}
