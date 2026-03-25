package parser

import "fmt"

const (
	// GridSize is the size of the tetromino grid (4x4)
	GridSize = 4
	// RequiredBlocks is the number of blocks each tetromino must have
	RequiredBlocks = 4
)

type point struct {
	row, column int
}

// ValidateGrid checks that the grid has exactly 4 '#' cells and they form a single connected component (4-directionally)
func ValidateGrid(grid [][]rune, tetrominoNumber int) error {
	var cells []point

	// Count all '#' characters
	for row := 0; row < GridSize; row++ {
		for column := 0; column < GridSize; column++ {
			if grid[row][column] == '#' {
				cells = append(cells, point{row, column})
			}
		}
	}

	// Check block count
	if len(cells) != RequiredBlocks {
		return fmt.Errorf("Tetromino #%d: Invalid number of blocks - expected %d '#' characters, got %d", tetrominoNumber, RequiredBlocks, len(cells))
	}

	// Check connectivity using BFS (Breadth-First Search)
	visited := map[point]bool{}
	queue := []point{cells[0]}
	visited[cells[0]] = true

	// Four directions: down, up, right, left
	directions := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, direction := range directions {
			neighbor := point{current.row + direction.row, current.column + direction.column}

			// Check if neighbor is within grid bounds
			if neighbor.row < 0 || neighbor.row >= GridSize || neighbor.column < 0 || neighbor.column >= GridSize {
				continue
			}

			// Check if neighbor is a block
			if grid[neighbor.row][neighbor.column] != '#' {
				continue
			}

			// Add unvisited neighbors to queue
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	// All blocks must be reachable from the first block
	if len(visited) != RequiredBlocks {
		return fmt.Errorf("Tetromino #%d: Blocks not connected - all %d blocks must be adjacent (up/down/left/right)", tetrominoNumber, RequiredBlocks)
	}

	return nil
}
