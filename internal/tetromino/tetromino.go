package tetromino

// Point represents a coordinate on the board or within a tetromino
type Point struct {
	Row, Column int
}

// Tetromino represents a tetris piece with its block positions and display ID
type Tetromino struct {
	Blocks []Point // Normalized block coordinates (relative to origin)
	ID     rune    // Display character (A, B, C, ...)
}

// FromGrid builds a Tetromino from a 4x4 grid of '.' and '#' characters and normalizes it
func FromGrid(grid [][]rune, id rune) Tetromino {
	var points []Point
	for row := range grid {
		for column, character := range grid[row] {
			if character == '#' {
				points = append(points, Point{row, column})
			}
		}
	}
	points = Normalize(points)
	return Tetromino{Blocks: points, ID: id}
}
