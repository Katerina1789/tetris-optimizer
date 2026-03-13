package board_test

import (
	"testing"
	"tetris-optimizer/internal/board"
)

func TestNew(t *testing.T) {
	b := board.New(5)
	if b.Size != 5 {
		t.Errorf("expected size 5, got %d", b.Size)
	}
	for r := 0; r < b.Size; r++ {
		for c := 0; c < b.Size; c++ {
			if b.Cells[r][c] != '.' {
				t.Errorf("expected '.', got %c at (%d,%d)", b.Cells[r][c], r, c)
			}
		}
	}
}

func TestString(t *testing.T) {
	b := board.New(2)
	b.Cells[0][0] = 'A'
	b.Cells[1][1] = 'B'
	expected := "A.\n.B\n"
	if b.String() != expected {
		t.Errorf("expected %q, got %q", expected, b.String())
	}
}

func TestInBounds(t *testing.T) {
	b := board.New(3)
	tests := []struct {
		r, c int
		want bool
	}{
		{0, 0, true},
		{2, 2, true},
		{-1, 0, false},
		{0, -1, false},
		{3, 0, false},
		{0, 3, false},
	}
	for _, tt := range tests {
		if got := b.InBounds(tt.r, tt.c); got != tt.want {
			t.Errorf("InBounds(%d,%d) = %v, want %v", tt.r, tt.c, got, tt.want)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	b := board.New(2)
	if !b.IsEmpty(0, 0) {
		t.Error("expected empty cell")
	}
	b.Cells[0][0] = 'A'
	if b.IsEmpty(0, 0) {
		t.Error("expected non-empty cell")
	}
}
