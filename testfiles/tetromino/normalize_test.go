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

	result := tetromino.FromGrid(grid, 'A')
	if result.Blocks[0].Row != 0 || result.Blocks[0].Column != 0 {
		t.Errorf("expected first block at (0,0), got (%d,%d)", result.Blocks[0].Row, result.Blocks[0].Column)
	}
}

func TestNormalizeOffset(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.'},
		{'.', '.', '#', '#'},
		{'.', '.', '#', '#'},
		{'.', '.', '.', '.'},
	}

	result := tetromino.FromGrid(grid, 'B')
	minimumRow, minimumColumn := result.Blocks[0].Row, result.Blocks[0].Column
	for _, point := range result.Blocks {
		if point.Row < minimumRow {
			minimumRow = point.Row
		}
		if point.Column < minimumColumn {
			minimumColumn = point.Column
		}
	}

	if minimumRow != 0 || minimumColumn != 0 {
		t.Errorf("expected normalized to start at (0,0), got min (%d,%d)", minimumRow, minimumColumn)
	}
}
