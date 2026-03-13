package board_test

import (
	"testing"
	"tetris-optimizer/internal/board"
	"tetris-optimizer/internal/tetromino"
)

func TestPlace(t *testing.T) {
	b := board.New(4)
	tet := tetromino.Tetromino{
		Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 0, C: 1}, {R: 1, C: 0}, {R: 1, C: 1}},
		ID:     'A',
	}
	if !b.Place(tet, 0, 0) {
		t.Error("expected successful placement")
	}
	if b.Cells[0][0] != 'A' || b.Cells[0][1] != 'A' || b.Cells[1][0] != 'A' || b.Cells[1][1] != 'A' {
		t.Error("tetromino not placed correctly")
	}
}

func TestPlaceOutOfBounds(t *testing.T) {
	b := board.New(4)
	tet := tetromino.Tetromino{
		Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 0, C: 1}, {R: 1, C: 0}, {R: 1, C: 1}},
		ID:     'A',
	}
	if b.Place(tet, 3, 3) {
		t.Error("expected placement to fail (out of bounds)")
	}
}

func TestPlaceOverlap(t *testing.T) {
	b := board.New(4)
	tet1 := tetromino.Tetromino{
		Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 0, C: 1}},
		ID:     'A',
	}
	tet2 := tetromino.Tetromino{
		Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 1, C: 0}},
		ID:     'B',
	}
	b.Place(tet1, 0, 0)
	if b.Place(tet2, 0, 0) {
		t.Error("expected placement to fail (overlap)")
	}
}

func TestRemove(t *testing.T) {
	b := board.New(4)
	tet := tetromino.Tetromino{
		Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 0, C: 1}, {R: 1, C: 0}, {R: 1, C: 1}},
		ID:     'A',
	}
	b.Place(tet, 0, 0)
	b.Remove(tet, 0, 0)
	if !b.IsEmpty(0, 0) || !b.IsEmpty(0, 1) || !b.IsEmpty(1, 0) || !b.IsEmpty(1, 1) {
		t.Error("tetromino not removed correctly")
	}
}

func TestFirstEmpty(t *testing.T) {
	b := board.New(3)
	r, c, found := b.FirstEmpty()
	if !found || r != 0 || c != 0 {
		t.Errorf("expected first empty at (0,0), got (%d,%d) found=%v", r, c, found)
	}
	
	b.Cells[0][0] = 'A'
	r, c, found = b.FirstEmpty()
	if !found || r != 0 || c != 1 {
		t.Errorf("expected first empty at (0,1), got (%d,%d) found=%v", r, c, found)
	}
	
	// Fill entire board
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.Cells[i][j] = 'X'
		}
	}
	r, c, found = b.FirstEmpty()
	if found {
		t.Error("expected no empty cells found")
	}
}
