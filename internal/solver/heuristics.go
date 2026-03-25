package solver

import "tetris-optimizer/internal/tetromino"

// SortByConstraint orders tetrominoes so the most constrained ones are placed first.
// Constraint = fewer unique rotations → fewer valid placements → better pruning.
// This heuristic improves backtracking performance by reducing the search space early.
//
// Example: A square has 1 unique rotation (most constrained), while an L-shape has 4 rotations.
// Placing the square first eliminates more possibilities, making subsequent placements faster.
//
// Returns the precomputed rotation sets for each tetromino (already sorted).
func SortByConstraint(tetrominoes []tetromino.Tetromino) [][]tetromino.Tetromino {
	// Create pairs of (tetromino, rotations) for sorting
	type tetrominoPair struct {
		tetromino tetromino.Tetromino
		rotations []tetromino.Tetromino
	}

	// Calculate rotations once for each tetromino
	pairs := make([]tetrominoPair, len(tetrominoes))
	for index, currentTetromino := range tetrominoes {
		rotations := tetromino.Rotations(currentTetromino)
		pairs[index] = tetrominoPair{
			tetromino: currentTetromino,
			rotations: rotations,
		}
	}

	// Sort using insertion sort (simple, stable, good for small arrays)
	// Sorts in ascending order: fewer rotations first (most constrained)
	for outerIndex := 1; outerIndex < len(pairs); outerIndex++ {
		innerIndex := outerIndex
		for innerIndex > 0 && len(pairs[innerIndex].rotations) < len(pairs[innerIndex-1].rotations) {
			pairs[innerIndex], pairs[innerIndex-1] = pairs[innerIndex-1], pairs[innerIndex]
			innerIndex--
		}
	}

	// Extract sorted tetrominoes and rotation sets
	rotationSets := make([][]tetromino.Tetromino, len(tetrominoes))
	for index := range tetrominoes {
		tetrominoes[index] = pairs[index].tetromino
		rotationSets[index] = pairs[index].rotations
	}

	return rotationSets
}
