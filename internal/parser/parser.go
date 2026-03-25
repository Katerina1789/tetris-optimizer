package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"tetris-optimizer/internal/tetromino"
)

// ParseFile reads the file at path and returns a slice of tetrominoes labeled A, B, C, ... in the order they appear.
func ParseFile(filePath string) ([]tetromino.Tetromino, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Cannot open file: %v", err)
	}
	defer file.Close()

	var blocks [][]string
	var currentBlock []string
	var allErrors []string

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			if len(currentBlock) > 0 {
				blocks = append(blocks, currentBlock)
				currentBlock = nil
			}
			continue
		}
		currentBlock = append(currentBlock, line)
	}

	// Add last block if exists
	if len(currentBlock) > 0 {
		blocks = append(blocks, currentBlock)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}

	// Parse each tetromino block and collect all errors
	var tetrominoes []tetromino.Tetromino
	for blockIndex, block := range blocks {
		tetrominoNumber := blockIndex + 1
		tetrominoID := rune('A' + blockIndex)

		parsedTetromino, err := ParseTetrominoBlock(block, tetrominoID, tetrominoNumber)
		if err != nil {
			allErrors = append(allErrors, err.Error())
			continue // Collect all errors instead of stopping at first
		}
		tetrominoes = append(tetrominoes, parsedTetromino)
	}

	// If there were any errors, return them all
	if len(allErrors) > 0 {
		errorMessage := strings.Join(allErrors, "\n")
		return nil, fmt.Errorf("%s", errorMessage)
	}

	return tetrominoes, nil
}

// ParseTetrominoBlock converts a 4x4 block of lines into a validated Tetromino
func ParseTetrominoBlock(lines []string, id rune, tetrominoNumber int) (tetromino.Tetromino, error) {
	if len(lines) != 4 {
		return tetromino.Tetromino{}, fmt.Errorf("Tetromino #%d: Invalid height - expected 4 lines, got %d", tetrominoNumber, len(lines))
	}

	grid := make([][]rune, 4)
	for lineIndex := 0; lineIndex < 4; lineIndex++ {
		if len(lines[lineIndex]) != 4 {
			return tetromino.Tetromino{}, fmt.Errorf("Tetromino #%d, line %d: Invalid width - expected 4 characters, got %d", tetrominoNumber, lineIndex+1, len(lines[lineIndex]))
		}

		row := []rune(lines[lineIndex])
		for _, character := range row {
			if character != '#' && character != '.' {
				return tetromino.Tetromino{}, fmt.Errorf("Tetromino #%d: Invalid character '%c' - only '#' and '.' are allowed", tetrominoNumber, character)
			}
		}
		grid[lineIndex] = row
	}

	if err := ValidateGrid(grid, tetrominoNumber); err != nil {
		return tetromino.Tetromino{}, err
	}

	return tetromino.FromGrid(grid, id), nil
}
