package board_test

import (
	"testing"

	"tetris-optimizer/internal/board"
	"tetris-optimizer/internal/tetromino"
)

func TestNew(t *testing.T) {
	testBoard := board.New(5)
	if testBoard.Size != 5 {
		t.Errorf("expected size 5, got %d", testBoard.Size)
	}
	for row := 0; row < testBoard.Size; row++ {
		for column := 0; column < testBoard.Size; column++ {
			if testBoard.Cell(row, column) != '.' {
				t.Errorf("expected '.', got %c at (%d,%d)", testBoard.Cell(row, column), row, column)
			}
		}
	}
}

func TestString(t *testing.T) {
	testBoard := board.New(2)
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
	expected := "AA\nAA\n"
	if testBoard.String() != expected {
		t.Errorf("expected %q, got %q", expected, testBoard.String())
	}
}

func TestInBounds(t *testing.T) {
	testBoard := board.New(3)
	tests := []struct {
		row, column int
		want        bool
	}{
		{0, 0, true},
		{2, 2, true},
		{-1, 0, false},
		{0, -1, false},
		{3, 0, false},
		{0, 3, false},
	}
	for _, testCase := range tests {
		if got := testBoard.InBounds(testCase.row, testCase.column); got != testCase.want {
			t.Errorf("InBounds(%d,%d) = %v, want %v", testCase.row, testCase.column, got, testCase.want)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	testBoard := board.New(2)
	if !testBoard.IsEmpty(0, 0) {
		t.Error("expected empty cell")
	}
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
	if testBoard.IsEmpty(0, 0) {
		t.Error("expected non-empty cell")
	}
}
