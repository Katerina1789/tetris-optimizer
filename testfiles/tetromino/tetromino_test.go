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

	result := tetromino.FromGrid(grid, 'A')

	if result.ID != 'A' {
		t.Fatalf("expected ID A, got %c", result.ID)
	}
	if len(result.Blocks) != 4 {
		t.Fatalf("expected 4 blocks, got %d", len(result.Blocks))
	}
}
