package tetromino

// normalize shifts blocks so that the shape's top-left is at (0,0).
func normalize(blocks []Point) []Point {
	if len(blocks) == 0 {
		return blocks
	}
	minR, minC := blocks[0].R, blocks[0].C
	for _, p := range blocks {
		if p.R < minR {
			minR = p.R
		}
		if p.C < minC {
			minC = p.C
		}
	}
	out := make([]Point, len(blocks))
	for i, p := range blocks {
		out[i] = Point{p.R - minR, p.C - minC}
	}
	return out
}
