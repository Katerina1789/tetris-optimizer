package tetromino

// Rotations returns all unique 90-degree rotations of a tetromino.
func Rotations(t Tetromino) []Tetromino {
	var res []Tetromino
	seen := map[string]bool{}

	cur := t.Blocks
	for i := 0; i < 4; i++ {
		norm := Normalize(cur)
		key := keyOf(norm)
		if !seen[key] {
			seen[key] = true
			res = append(res, Tetromino{Blocks: norm, ID: t.ID})
		}
		cur = rotate90(cur)
	}
	return res
}

// rotate90 rotates points 90 degrees around origin.
func rotate90(blocks []Point) []Point {
	out := make([]Point, len(blocks))
	for i, p := range blocks {
		out[i] = Point{R: p.C, C: -p.R}
	}
	return out
}

// keyOf builds a compact key for a set of points to detect duplicates.
func keyOf(blocks []Point) string {
	b := make([]byte, 0, len(blocks)*4)
	for _, p := range blocks {
		b = append(b, byte(p.R+16), byte(p.C+16))
	}
	return string(b)
}
