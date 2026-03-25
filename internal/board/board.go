package board

import (
	"strings"
)

// Board represents a square game board for placing tetrominoes
type Board struct {
	Size  int
	cells [][]rune // Private: use Cell() to access
}

// New creates a Size x Size board filled with '.' (empty cells)
func New(size int) *Board {
	cells := make([][]rune, size)
	for row := range cells {
		cells[row] = make([]rune, size)
		for column := range cells[row] {
			cells[row][column] = '.'
		}
	}
	return &Board{Size: size, cells: cells}
}

// String returns the board as lines of text with newlines
func (board *Board) String() string {
	var stringBuilder strings.Builder
	for row := 0; row < board.Size; row++ {
		for column := 0; column < board.Size; column++ {
			stringBuilder.WriteRune(board.cells[row][column])
		}
		stringBuilder.WriteRune('\n')
	}
	return stringBuilder.String()
}

// InBounds checks if the given coordinates are inside the board
func (board *Board) InBounds(row, column int) bool {
	return row >= 0 && row < board.Size && column >= 0 && column < board.Size
}

// IsEmpty reports whether the cell at (row, column) is empty ('.')
func (board *Board) IsEmpty(row, column int) bool {
	return board.cells[row][column] == '.'
}

// Cell returns the rune at the given position (for reading)
func (board *Board) Cell(row, column int) rune {
	return board.cells[row][column]
}

// setCell sets the rune at the given position (internal use only)
func (board *Board) setCell(row, column int, value rune) {
	board.cells[row][column] = value
}
