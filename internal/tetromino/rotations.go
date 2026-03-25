package tetromino

const (
	// KeyOffset is added to coordinates when creating string keys to handle negative values
	// Supports coordinates from -16 to 239
	KeyOffset = 16
)

// Rotations returns all unique 90-degree rotations of a tetromino
func Rotations(tetrominoToRotate Tetromino) []Tetromino {
	var results []Tetromino
	seenKeys := map[string]bool{}

	currentBlocks := tetrominoToRotate.Blocks
	for rotation := 0; rotation < 4; rotation++ {
		normalized := Normalize(currentBlocks)
		key := keyOf(normalized)
		if !seenKeys[key] {
			seenKeys[key] = true
			results = append(results, Tetromino{Blocks: normalized, ID: tetrominoToRotate.ID})
		}
		currentBlocks = rotate90(currentBlocks)
	}

	return results
}

// rotate90 rotates points 90 degrees clockwise around the origin
// Formula: (row, column) -> (column, -row)
func rotate90(blocks []Point) []Point {
	rotated := make([]Point, len(blocks))
	for index, point := range blocks {
		rotated[index] = Point{Row: point.Column, Column: -point.Row}
	}
	return rotated
}

// keyOf builds a compact string key for a set of points to detect duplicate rotations
func keyOf(blocks []Point) string {
	bytes := make([]byte, 0, len(blocks)*2)
	for _, point := range blocks {
		bytes = append(bytes, byte(point.Row+KeyOffset), byte(point.Column+KeyOffset))
	}
	return string(bytes)
}
