# Project Structure

## Repository Layout

```
tetris-optimizer/
├── cmd/tetris-optimizer/
│   └── main.go                  # Entry point
├── internal/
│   ├── board/
│   │   ├── board.go             # Board structure
│   │   └── place.go             # Placement operations
│   ├── parser/
│   │   ├── parser.go            # File parsing
│   │   └── validate.go          # Input validation
│   ├── solver/
│   │   ├── solver.go            # Main solver
│   │   ├── backtracking.go      # Backtracking algorithm
│   │   └── heuristics.go        # Constraint heuristic
│   └── tetromino/
│       ├── tetromino.go         # Tetromino structure
│       ├── rotations.go         # Rotation generation
│       └── normalize.go         # Coordinate normalization
├── testfiles/
│   ├── board/
│   │   ├── board_test.go        # Board tests
│   │   └── place_test.go        # Placement tests
│   ├── parser/
│   │   ├── parser_test.go       # Parser tests
│   │   └── validate_test.go     # Validation tests
│   ├── solver/
│   │   ├── solver_test.go       # Solver tests
│   │   ├── backtracking_test.go # Backtracking tests
│   │   └── heuristics_test.go   # Heuristic tests
│   └── tetromino/
│       ├── tetromino_test.go    # Tetromino tests
│       ├── rotations_test.go    # Rotation tests
│       └── normalize_test.go    # Normalization tests
├── audit/
│   ├── examples/*.txt           # Test input files (optional)
│   └── audit_guide.md           # Audit instructions
├── docs/
│   ├── project_structure.md     # This file
│   ├── architecture.md          # System design
│   └── algorithm.md             # Algorithm explanation
├── .gitignore                   # Git exclusions
├── go.mod                       # Go module
├── LICENSE                      # MIT License
├── Makefile                     # Build automation
└── README.md                    # Main documentation
```

---

## Source Code (11 files, 525 lines)

### cmd/tetris-optimizer/main.go (35 lines)
Entry point. Validates arguments, parses input file, runs solver, prints result or ERROR. Uses exit code 1 on errors.

### internal/board/board.go (55 lines)
Board structure with Size and private cells. Functions: New (creates board), String (formats output), InBounds, IsEmpty, Cell (getter), setCell (private setter).

### internal/board/place.go (49 lines)
Placement operations. Place (two-phase validation: check all blocks, then place all), Remove (clears tetromino), FirstEmpty (finds first empty cell for backtracking).

### internal/parser/parser.go (99 lines)
File parsing. ParseFile reads file, splits into 4-line blocks, validates each block, creates tetrominoes with IDs A-Z. Collects all errors with line numbers.

### internal/parser/validate.go (73 lines)
Input validation. ValidateGrid uses BFS to check 4 connected blocks. Constants: GridSize=4, RequiredBlocks=4.

### internal/solver/solver.go (47 lines)
Main solver. Validates input (empty check, max 26), applies heuristic, calculates min size = ceil(sqrt(n×4)), tries increasing board sizes. Constants: MaxTetrominoes=26, MaxBoardSize=20.

### internal/solver/backtracking.go (58 lines)
Backtracking algorithm. Finds first empty cell, tries all rotations at positions -3 to +3 relative to empty cell, recursively places remaining pieces. Constant: MaxTetrominoSize=3.

### internal/solver/heuristics.go (44 lines)
Constraint-based sorting. SortByConstraint sorts tetrominoes by rotation count (fewer first), precomputes rotation sets once. Most Constrained Variable heuristic reduces backtracking.

### internal/tetromino/tetromino.go (26 lines)
Tetromino structure. Types: Point (Row, Column), Tetromino (Blocks, ID). FromGrid extracts blocks from grid and normalizes to origin.

### internal/tetromino/rotations.go (45 lines)
Rotation generation. Rotations generates 1-4 unique 90° rotations using formula (row,column)→(column,-row). Deduplicates using string keys. Constant: KeyOffset=16.

### internal/tetromino/normalize.go (27 lines)
Coordinate normalization. Normalize shifts all blocks so minimum row and column are 0. Enables consistent representation and rotation comparison.

---

## Test Files (10 files, 35 tests, 94.8% coverage)

### testfiles/board/board_test.go (9 tests)
Tests board creation, string formatting, boundary checking, empty cell detection. Coverage: 100% of board.go.

### testfiles/board/place_test.go (9 tests)
Tests placement (success, out-of-bounds, overlap), removal, first empty cell finding. Coverage: 100% of place.go.

### testfiles/parser/parser_test.go (5 tests)
Tests file parsing (single, multiple, invalid height/width/characters). Coverage: 94.3% of parser.go.

### testfiles/parser/validate_test.go (3 tests)
Tests grid validation (valid, wrong block count, disconnected blocks). Coverage: 100% of validate.go.

### testfiles/solver/solver_test.go (4 tests)
Tests solving (single piece, min size calculation, empty input, too many pieces). Coverage: 91.7% of solver.go.

### testfiles/solver/backtracking_test.go (1 test)
Tests backtracking with multiple tetrominoes. Coverage: 81.2% of backtracking.go.

### testfiles/solver/heuristics_test.go (5 tests)
Tests constraint sorting (basic, empty, single, stability, actual rotations). Coverage: 86.7% of heuristics.go.

### testfiles/tetromino/tetromino_test.go (1 test)
Tests tetromino creation from grid with normalization. Coverage: 100% of tetromino.go.

### testfiles/tetromino/rotations_test.go (5 tests)
Tests rotation generation (square=1, line=2, L/T=4, uniqueness). Coverage: 100% of rotations.go.

### testfiles/tetromino/normalize_test.go (2 tests)
Tests coordinate normalization (already normalized, offset). Coverage: 91.7% of normalize.go.

---

## Documentation (3 files)

### docs/project_structure.md (this file)
Complete file listing with explanations of what each file does and how it works.

### docs/architecture.md (285 lines)
System design, component interactions, heuristics explanation (MCV), execution flows, performance characteristics, example walkthrough.

### docs/algorithm.md (184 lines)
Algorithm steps, key optimizations, complexity analysis, example walkthrough, performance results.

---

## Configuration (5 files)

### .gitignore
Excludes binaries, test files, coverage reports, profiling data, IDE files, OS files.

### go.mod
Go module definition. Version 1.25.7, no external dependencies (standard library only).

### LICENSE
MIT License. Copyright 2026 Katerina Kasdanastasi.

### Makefile
Build automation. Targets: build, run, test, coverage, bench, profile, clean, all.

### README.md
Main documentation. Overview, features, usage instructions, input/output formats, algorithm summary, testing guide.

---

## Audit Files (optional)

### audit/audit_guide.md
Comprehensive audit guide with 22 test cases, quick test commands, and evaluation checklist.

### audit/examples/*.txt (12 files)
Test input files for validation:
- **bad_example_00-04.txt, bad_format.txt**: Invalid inputs (should print ERROR)
- **good_example_00-03.txt**: Valid inputs with expected empty space counts
- **hard_example.txt**: Performance test (12 tetrominoes, < 30s)
- **test_26.txt**: Maximum input test (26 tetrominoes, ~5s)

---

## Summary

**Total:** 42 files  
**Source:** 11 files, 525 lines  
**Tests:** 10 files, 35 tests, 94.8% coverage  
**Docs:** 3 files  
**Config:** 5 files  
**Audit:** 13 files
