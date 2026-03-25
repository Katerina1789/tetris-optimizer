package tetromino_test

import (
	"testing"

	"tetris-optimizer/internal/tetromino"
)

func TestRotationsSquare(t *testing.T) {
	squareTetromino := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 0, Column: 1},
			{Row: 1, Column: 0},
			{Row: 1, Column: 1},
		},
		ID: 'A',
	}

	rotations := tetromino.Rotations(squareTetromino)
	if len(rotations) < 1 || len(rotations) > 4 {
		t.Errorf("square should have 1-4 rotations, got %d", len(rotations))
	}

	// All rotations should have same ID
	for _, rotation := range rotations {
		if rotation.ID != 'A' {
			t.Errorf("expected ID 'A' in rotation, got %c", rotation.ID)
		}
		if len(rotation.Blocks) != 4 {
			t.Errorf("expected 4 blocks in rotation, got %d", len(rotation.Blocks))
		}
	}
}

func TestRotationsLine(t *testing.T) {
	lineTetromino := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 1, Column: 0},
			{Row: 2, Column: 0},
			{Row: 3, Column: 0},
		},
		ID: 'B',
	}

	rotations := tetromino.Rotations(lineTetromino)
	if len(rotations) < 2 || len(rotations) > 4 {
		t.Errorf("line should have 2-4 rotations, got %d", len(rotations))
	}

	for _, rotation := range rotations {
		if rotation.ID != 'B' {
			t.Errorf("expected ID 'B' in rotation, got %c", rotation.ID)
		}
	}
}

func TestRotationsL(t *testing.T) {
	lShapeTetromino := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 1, Column: 0},
			{Row: 2, Column: 0},
			{Row: 2, Column: 1},
		},
		ID: 'C',
	}

	rotations := tetromino.Rotations(lShapeTetromino)
	if len(rotations) != 4 {
		t.Errorf("L-shape should have 4 unique rotations, got %d", len(rotations))
	}

	for _, rotation := range rotations {
		if rotation.ID != 'C' {
			t.Errorf("expected ID 'C' in rotation, got %c", rotation.ID)
		}
	}
}

func TestRotationsT(t *testing.T) {
	tShapeTetromino := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 1},
			{Row: 1, Column: 0},
			{Row: 1, Column: 1},
			{Row: 1, Column: 2},
		},
		ID: 'D',
	}

	rotations := tetromino.Rotations(tShapeTetromino)
	if len(rotations) != 4 {
		t.Errorf("T-shape should have 4 unique rotations, got %d", len(rotations))
	}

	for _, rotation := range rotations {
		if rotation.ID != 'D' {
			t.Errorf("expected ID 'D' in rotation, got %c", rotation.ID)
		}
	}
}

func TestRotationsUnique(t *testing.T) {
	// Test that rotations are actually unique
	lShapeTetromino := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 1, Column: 0},
			{Row: 2, Column: 0},
			{Row: 2, Column: 1},
		},
		ID: 'L',
	}

	rotations := tetromino.Rotations(lShapeTetromino)

	// Check that no two rotations are identical
	for firstIndex := 0; firstIndex < len(rotations); firstIndex++ {
		for secondIndex := firstIndex + 1; secondIndex < len(rotations); secondIndex++ {
			if blocksEqual(rotations[firstIndex].Blocks, rotations[secondIndex].Blocks) {
				t.Errorf("found duplicate rotations at indices %d and %d", firstIndex, secondIndex)
			}
		}
	}
}

func blocksEqual(firstBlocks, secondBlocks []tetromino.Point) bool {
	if len(firstBlocks) != len(secondBlocks) {
		return false
	}
	for index := range firstBlocks {
		if firstBlocks[index] != secondBlocks[index] {
			return false
		}
	}
	return true
}
