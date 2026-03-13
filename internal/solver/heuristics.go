package solver

import "tetris-optimizer/internal/tetromino"

// SortByConstraint orders tetrominoes so the most restrictive ones go first, restrictive = fewer unique rotations → fewer valid placements → better pruning.
func SortByConstraint(tetros []tetromino.Tetromino) {
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
