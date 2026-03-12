package board

// Package board defines the square board and basic operations on it.

import (
	"strings"
)

type Board struct {
	Size  int
	Cells [][]rune
}

// New creates a Size x Size board filled with '.'.
func New(size int) *Board {
	cells := make([][]rune, size)
	for i := range cells {
		cells[i] = make([]rune, size)
		for j := range cells[i] {
			cells[i][j] = '.'
		}
	}
	return &Board{Size: size, Cells: cells}
}

// String returns the board as lines of text.
func (b *Board) String() string {
	var sb strings.Builder
	for r := 0; r < b.Size; r++ {
		for c := 0; c < b.Size; c++ {
			sb.WriteRune(b.Cells[r][c])
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

// InBounds checks if (r,c) is inside the board.
func (b *Board) InBounds(r, c int) bool {
	return r >= 0 && r < b.Size && c >= 0 && c < b.Size
}

// IsEmpty reports whether cell (r,c) is '.'.
func (b *Board) IsEmpty(r, c int) bool {
	return b.Cells[r][c] == '.'
}
