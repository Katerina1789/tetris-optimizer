package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"tetris-optimizer/internal/tetromino"
)

// ParseFile reads the file at path and returns a slice of tetrominoes labeled A, B, C, ... in the order they appear.
func ParseFile(path string) ([]tetromino.Tetromino, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Cannot open file: %v", err)
	}
	defer f.Close()

	var blocks [][]string
	var current []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			if len(current) > 0 {
				blocks = append(blocks, current)
				current = nil
			}
			continue
		}
		current = append(current, line)
	}
	if len(current) > 0 {
		blocks = append(blocks, current)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var tetros []tetromino.Tetromino
	for i, b := range blocks {
		t, err := ParseTetrominoBlock(b, rune('A'+i))
		if err != nil {
			return nil, err
		}
		tetros = append(tetros, t)
	}
	return tetros, nil
}

// ParseTetrominoBlock converts a 4x4 block of lines into a validated Tetromino.
func ParseTetrominoBlock(lines []string, id rune) (tetromino.Tetromino, error) {
	if len(lines) != 4 {
		return tetromino.Tetromino{}, fmt.Errorf("Invalid tetromino height: expected 4 lines, got %d", len(lines))
	}
	grid := make([][]rune, 4)
	for i := 0; i < 4; i++ {
		if len(lines[i]) != 4 {
			return tetromino.Tetromino{}, fmt.Errorf("Invalid tetromino width: expected 4 characters per line, got %d on line %d", len(lines[i]), i+1)
		}
		row := []rune(lines[i])
		for _, c := range row {
			if c != '#' && c != '.' {
				return tetromino.Tetromino{}, fmt.Errorf("Invalid character '%c': only '#' and '.' are allowed", c)
			}
		}
		grid[i] = row
	}
	if err := ValidateGrid(grid); err != nil {
		return tetromino.Tetromino{}, err
	}
	return tetromino.FromGrid(grid, id), nil
}
