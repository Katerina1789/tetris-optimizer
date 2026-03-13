package tetromino_test

import (
	"testing"

	"tetris-optimizer/internal/tetromino"
)

func TestRotationsSquare(t *testing.T) {
	tet := tetromino.Tetromino{
		Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 0, C: 1}, {R: 1, C: 0}, {R: 1, C: 1}},
		ID:     'A',
	}
	rots := tetromino.Rotations(tet)
	if len(rots) < 1 || len(rots) > 4 {
		t.Errorf("square should have 1-4 rotations, got %d", len(rots))
	}
	// All rotations should have same ID
	for _, rot := range rots {
		if rot.ID != 'A' {
			t.Errorf("expected ID 'A' in rotation, got %c", rot.ID)
		}
		if len(rot.Blocks) != 4 {
			t.Errorf("expected 4 blocks in rotation, got %d", len(rot.Blocks))
		}
	}
}

func TestRotationsLine(t *testing.T) {
	tet := tetromino.Tetromino{
		Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 1, C: 0}, {R: 2, C: 0}, {R: 3, C: 0}},
		ID:     'B',
	}
	rots := tetromino.Rotations(tet)
	if len(rots) < 2 || len(rots) > 4 {
		t.Errorf("line should have 2-4 rotations, got %d", len(rots))
	}
	for _, rot := range rots {
		if rot.ID != 'B' {
			t.Errorf("expected ID 'B' in rotation, got %c", rot.ID)
		}
	}
}

func TestRotationsL(t *testing.T) {
	tet := tetromino.Tetromino{
		Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 1, C: 0}, {R: 2, C: 0}, {R: 2, C: 1}},
		ID:     'C',
	}
	rots := tetromino.Rotations(tet)
	if len(rots) != 4 {
		t.Errorf("L-shape should have 4 unique rotations, got %d", len(rots))
	}
	for _, rot := range rots {
		if rot.ID != 'C' {
			t.Errorf("expected ID 'C' in rotation, got %c", rot.ID)
		}
	}
}

func TestRotationsT(t *testing.T) {
	tet := tetromino.Tetromino{
		Blocks: []tetromino.Point{{R: 0, C: 1}, {R: 1, C: 0}, {R: 1, C: 1}, {R: 1, C: 2}},
		ID:     'D',
	}
	rots := tetromino.Rotations(tet)
	if len(rots) != 4 {
		t.Errorf("T-shape should have 4 unique rotations, got %d", len(rots))
	}
	for _, rot := range rots {
		if rot.ID != 'D' {
			t.Errorf("expected ID 'D' in rotation, got %c", rot.ID)
		}
	}
}

func TestRotationsUnique(t *testing.T) {
	// Test that rotations are actually unique
	tet := tetromino.Tetromino{
		Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 1, C: 0}, {R: 2, C: 0}, {R: 2, C: 1}},
		ID:     'L',
	}
	rots := tetromino.Rotations(tet)
	
	// Check that no two rotations are identical
	for i := 0; i < len(rots); i++ {
		for j := i + 1; j < len(rots); j++ {
			if blocksEqual(rots[i].Blocks, rots[j].Blocks) {
				t.Errorf("found duplicate rotations at indices %d and %d", i, j)
			}
		}
	}
}

func blocksEqual(a, b []tetromino.Point) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
