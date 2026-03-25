package board_test

import (
	"testing"

	"tetris-optimizer/internal/board"
	"tetris-optimizer/internal/tetromino"
)

func TestPlace(t *testing.T) {
	testBoard := board.New(4)
	squareTetromino := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 0, Column: 1},
			{Row: 1, Column: 0},
			{Row: 1, Column: 1},
		},
		ID: 'A',
	}

	if !testBoard.Place(squareTetromino, 0, 0) {
		t.Error("expected successful placement")
	}

	if testBoard.Cell(0, 0) != 'A' || testBoard.Cell(0, 1) != 'A' ||
		testBoard.Cell(1, 0) != 'A' || testBoard.Cell(1, 1) != 'A' {
		t.Error("tetromino not placed correctly")
	}
}

func TestPlaceOutOfBounds(t *testing.T) {
	testBoard := board.New(4)
	squareTetromino := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 0, Column: 1},
			{Row: 1, Column: 0},
			{Row: 1, Column: 1},
		},
		ID: 'A',
	}

	if testBoard.Place(squareTetromino, 3, 3) {
		t.Error("expected placement to fail (out of bounds)")
	}
}

func TestPlaceOverlap(t *testing.T) {
	testBoard := board.New(4)
	firstTetromino := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 0, Column: 1},
		},
		ID: 'A',
	}
	secondTetromino := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 1, Column: 0},
		},
		ID: 'B',
	}

	testBoard.Place(firstTetromino, 0, 0)
	if testBoard.Place(secondTetromino, 0, 0) {
		t.Error("expected placement to fail (overlap)")
	}
}

func TestRemove(t *testing.T) {
	testBoard := board.New(4)
	squareTetromino := tetromino.Tetromino{
		Blocks: []tetromino.Point{
			{Row: 0, Column: 0},
			{Row: 0, Column: 1},
			{Row: 1, Column: 0},
			{Row: 1, Column: 1},
		},
		ID: 'A',
	}

	testBoard.Place(squareTetromino, 0, 0)
	testBoard.Remove(squareTetromino, 0, 0)

	if !testBoard.IsEmpty(0, 0) || !testBoard.IsEmpty(0, 1) ||
		!testBoard.IsEmpty(1, 0) || !testBoard.IsEmpty(1, 1) {
		t.Error("tetromino not removed correctly")
	}
}

func TestFirstEmpty(t *testing.T) {
	testBoard := board.New(3)
	row, column, found := testBoard.FirstEmpty()
	if !found || row != 0 || column != 0 {
		t.Errorf("expected first empty at (0,0), got (%d,%d) found=%v", row, column, found)
	}

	// Place a tetromino at (0,0)
	testBoard.Place(tetromino.Tetromino{
		Blocks: []tetromino.Point{{Row: 0, Column: 0}},
		ID:     'A',
	}, 0, 0)

	row, column, found = testBoard.FirstEmpty()
	if !found || row != 0 || column != 1 {
		t.Errorf("expected first empty at (0,1), got (%d,%d) found=%v", row, column, found)
	}

	// Fill entire board
	for rowIndex := 0; rowIndex < 3; rowIndex++ {
		for columnIndex := 0; columnIndex < 3; columnIndex++ {
			testBoard.Place(tetromino.Tetromino{
				Blocks: []tetromino.Point{{Row: 0, Column: 0}},
				ID:     'X',
			}, rowIndex, columnIndex)
		}
	}

	row, column, found = testBoard.FirstEmpty()
	if found {
		t.Error("expected no empty cells found")
	}
}
