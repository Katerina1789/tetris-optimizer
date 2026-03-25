package solver_test

import (
	"testing"

	"tetris-optimizer/internal/solver"
	"tetris-optimizer/internal/tetromino"
)

func TestSolveSingleTetromino(t *testing.T) {
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
		t.Errorf("expected solution, got error: %v", err)
	}
	if resultBoard == nil {
		t.Error("expected board, got nil")
	}
	if resultBoard.Size != 2 {
		t.Errorf("expected board size 2, got %d", resultBoard.Size)
	}
}

func TestSolveMinimumBoardSize(t *testing.T) {
	// 1 tetromino = 4 blocks, min board = ceil(sqrt(4)) = 2
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
		t.Errorf("unexpected error: %v", err)
	}
	if resultBoard.Size != 2 {
		t.Errorf("expected minimum board size 2, got %d", resultBoard.Size)
	}
}

func TestSolveEmptyInput(t *testing.T) {
	tetrominoes := []tetromino.Tetromino{}

	_, err := solver.Solve(tetrominoes)
	if err == nil {
		t.Error("expected error for empty input")
	}
}

func TestSolveTooManyTetrominoes(t *testing.T) {
	// Create 27 tetrominoes (more than 26 = A-Z)
	tetrominoes := make([]tetromino.Tetromino, 27)
	for index := 0; index < 27; index++ {
		tetrominoes[index] = tetromino.Tetromino{
			Blocks: []tetromino.Point{
				{Row: 0, Column: 0},
				{Row: 0, Column: 1},
				{Row: 1, Column: 0},
				{Row: 1, Column: 1},
			},
			ID: rune('A' + index),
		}
	}

	_, err := solver.Solve(tetrominoes)
	if err == nil {
		t.Error("expected error for too many tetrominoes")
	}
}
