package solver_test

import (
	"testing"

	"tetris-optimizer/internal/solver"
	"tetris-optimizer/internal/tetromino"
)

func TestSolveSingleTetromino(t *testing.T) {
	tetros := []tetromino.Tetromino{
		{
			Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 0, C: 1}, {R: 1, C: 0}, {R: 1, C: 1}},
			ID:     'A',
		},
	}
	board, err := solver.Solve(tetros)
	if err != nil {
		t.Errorf("expected solution, got error: %v", err)
	}
	if board == nil {
		t.Error("expected board, got nil")
	}
	if board.Size != 2 {
		t.Errorf("expected board size 2, got %d", board.Size)
	}
}

func TestSolveMinimumBoardSize(t *testing.T) {
	// 1 tetromino = 4 blocks, min board = ceil(sqrt(4)) = 2
	tetros := []tetromino.Tetromino{
		{
			Blocks: []tetromino.Point{{R: 0, C: 0}, {R: 0, C: 1}, {R: 1, C: 0}, {R: 1, C: 1}},
			ID:     'A',
		},
	}
	board, err := solver.Solve(tetros)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if board.Size != 2 {
		t.Errorf("expected minimum board size 2, got %d", board.Size)
	}
}
