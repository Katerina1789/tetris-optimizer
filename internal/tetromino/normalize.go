package tetromino

// Normalize shifts all blocks so that the shape's top-left corner is at (0,0)
func Normalize(blocks []Point) []Point {
	if len(blocks) == 0 {
		return blocks
	}

	// Find minimum row and column values
	minimumRow, minimumColumn := blocks[0].Row, blocks[0].Column
	for _, point := range blocks {
		if point.Row < minimumRow {
			minimumRow = point.Row
		}
		if point.Column < minimumColumn {
			minimumColumn = point.Column
		}
	}

	// Shift all blocks to start at (0,0)
	normalized := make([]Point, len(blocks))
	for index, point := range blocks {
		normalized[index] = Point{point.Row - minimumRow, point.Column - minimumColumn}
	}

	return normalized
}
