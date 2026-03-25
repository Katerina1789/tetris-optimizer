package solver_test

import (
	"testing"

	"tetris-optimizer/internal/solver"
	"tetris-optimizer/internal/tetromino"
)

func TestBacktrackingSimple(t *testing.T) {
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

	resultBoard, err := solver.Solve(tetrominoes)
	if err != nil {
		t.Errorf("backtracking failed: %v", err)
	}
	if resultBoard == nil {
		t.Error("expected board from backtracking")
	}
}
