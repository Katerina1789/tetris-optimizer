package solver

// Heuristics for improving backtracking performance.
// These keep the solver minimal, clean, and audit‑friendly.

import "tetris-optimizer/internal/tetromino"

// sortByConstraint orders tetrominoes so the most restrictive ones go first.
// Restrictive = fewer unique rotations → fewer valid placements → better pruning.
func sortByConstraint(tetros []tetromino.Tetromino) {
	type pair struct {
		t   tetromino.Tetromino
		key int
	}

	arr := make([]pair, len(tetros))
	for i, t := range tetros {
		rots := tetromino.Rotations(t)
		arr[i] = pair{t, len(rots)}
	}

	// simple insertion sort (keeps code minimal, avoids importing "sort")
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j].key < arr[j-1].key {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}

	for i := range tetros {
		tetros[i] = arr[i].t
	}
}
