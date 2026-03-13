package parser

import "fmt"

type point struct{ r, c int }

// ValidateGrid checks that the grid has exactly 4 '#' cells and they form a single connected component (4-directionally).
func ValidateGrid(grid [][]rune) error {
	var cells []point
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if grid[r][c] == '#' {
				cells = append(cells, point{r, c})
			}
		}
	}
	if len(cells) != 4 {
		return fmt.Errorf("must have 4 blocks")
	}

	visited := map[point]bool{}
	var queue []point
	queue = append(queue, cells[0])
	visited[cells[0]] = true

	dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			n := point{cur.r + d.r, cur.c + d.c}
			if n.r < 0 || n.r >= 4 || n.c < 0 || n.c >= 4 {
				continue
			}
			if grid[n.r][n.c] != '#' {
				continue
			}
			if !visited[n] {
				visited[n] = true
				queue = append(queue, n)
			}
		}
	}
	if len(visited) != 4 {
		return fmt.Errorf("blocks not connected")
	}
	return nil
}
