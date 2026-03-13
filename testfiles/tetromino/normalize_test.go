package tetromino_test

import (
	"testing"

	"tetris-optimizer/internal/tetromino"
)

func TestNormalizeAlreadyNormalized(t *testing.T) {
	grid := [][]rune{
		{'#', '#', '.', '.'},
		{'#', '#', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
	}
	tet := tetromino.FromGrid(grid, 'A')
	if tet.Blocks[0].R != 0 || tet.Blocks[0].C != 0 {
		t.Errorf("expected first block at (0,0), got (%d,%d)", tet.Blocks[0].R, tet.Blocks[0].C)
	}
}

func TestNormalizeOffset(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.'},
		{'.', '.', '#', '#'},
		{'.', '.', '#', '#'},
		{'.', '.', '.', '.'},
	}
	tet := tetromino.FromGrid(grid, 'B')
	minR, minC := tet.Blocks[0].R, tet.Blocks[0].C
	for _, p := range tet.Blocks {
		if p.R < minR {
			minR = p.R
		}
		if p.C < minC {
			minC = p.C
		}
	}
	if minR != 0 || minC != 0 {
		t.Errorf("expected normalized to start at (0,0), got min (%d,%d)", minR, minC)
	}
}
