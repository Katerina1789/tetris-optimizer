package tetromino_test

import (
	"testing"
	"tetris-optimizer/internal/tetromino"
)

func TestFromGrid(t *testing.T) {
	grid := [][]rune{
		[]rune("#..."),
		[]rune("#..."),
		[]rune("#..."),
		[]rune("#..."),
	}

	tetro := tetromino.FromGrid(grid, 'A')

	if tetro.ID != 'A' {
		t.Fatalf("expected ID A, got %c", tetro.ID)
	}
	if len(tetro.Blocks) != 4 {
		t.Fatalf("expected 4 blocks, got %d", len(tetro.Blocks))
	}
}
