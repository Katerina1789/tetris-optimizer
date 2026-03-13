package solver_test

import (
	"testing"

	"tetris-optimizer/internal/solver"
	"tetris-optimizer/internal/tetromino"
)

func TestBacktrackingSimple(t *testing.T) {
	tetros := []tetromino.Tetromino{
		{
			Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 0, C: 1}, {R: 1, C: 0}, {R: 1, C: 1}},
			ID:     'A',
		},
	}
	board, err := solver.Solve(tetros)
	if err != nil {
		t.Errorf("backtracking failed: %v", err)
	}
	if board == nil {
		t.Error("expected board from backtracking")
	}
}
