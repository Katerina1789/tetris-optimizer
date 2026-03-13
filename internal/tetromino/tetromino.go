package tetromino

type Point struct {
	R, C int
}

type Tetromino struct {
	Blocks []Point // normalized block coordinates
	ID     rune    // display character (A, B, C, ...)
}

// FromGrid builds a Tetromino from a 4x4 grid of '.' and '#' and normalizes it.
func FromGrid(grid [][]rune, id rune) Tetromino {
	var pts []Point
	for r := range grid {
		for c, ch := range grid[r] {
			if ch == '#' {
				pts = append(pts, Point{r, c})
			}
		}
	}
	pts = Normalize(pts)
	return Tetromino{Blocks: pts, ID: id}
}
