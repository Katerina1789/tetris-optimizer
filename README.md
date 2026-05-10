# Tetris-Optimizer

[![Go Version](https://img.shields.io/badge/go-1.23+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-orange?style=for-the-badge&logo=opensourceinitiative&logoColor=white)](LICENSE)
[![Zone01](https://img.shields.io/badge/Zone01-Athens-blue?style=for-the-badge)](https://zone01.gr/gr/)

A Go program that solves the tetromino packing problem using optimized backtracking. Given a set of tetrominoes, it finds the smallest square board that can fit all pieces without gaps or overlaps. Project made as part of the Zone01 Athens curriculum.

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Repository Structure](#repository-structure)
- [How to Run](#how-to-run)
- [Requirements](#requirements)
- [Algorithm](#algorithm)
- [Input Format](#input-format)
- [Output Format](#output-format)
- [Testing](#testing)
- [Documentation](#documentation)
- [Author](#author)
- [License](#license)

## Description

Solves the tetromino packing problem: finds the smallest square board that fits all given tetrominoes without gaps or overlaps. Uses optimized backtracking with constraint satisfaction.

## Features

- Parses and validates tetromino input files
- Generates all unique rotations (90° increments)
- Finds minimal board size using backtracking
- First-empty-cell strategy for efficient solving
- Constraint-based heuristic (most constrained first)
- 35 unit tests with 94.8% coverage

## Repository Structure

```
tetris-optimizer/
├── cmd/
│   └── tetris-optimizer/
│       └── main.go              # Application entry point
├── internal/
│   ├── board/                   # Board operations
│   │   ├── board.go
│   │   └── place.go
│   ├── parser/                  # Input parsing and validation
│   │   ├── parser.go
│   │   └── validate.go
│   ├── solver/                  # Solving algorithm
│   │   ├── solver.go
│   │   ├── backtracking.go
│   │   └── heuristics.go
│   └── tetromino/               # Tetromino structure and operations
│       ├── tetromino.go
│       ├── rotations.go
│       └── normalize.go
├── testfiles/                   # Unit tests (35 tests, 94.8% coverage)
│   ├── board/                   # Board tests (9 tests)
│   ├── parser/                  # Parser tests (8 tests)
│   ├── solver/                  # Solver tests (10 tests)
│   └── tetromino/               # Tetromino tests (8 tests)
├── audit/
│   ├── examples/                # Test cases
│   │   ├── bad_example_*.txt   # Invalid inputs
│   │   ├── good_example_*.txt  # Valid inputs
│   │   └── hard_example.txt    # Performance test
│   └── audit_guide.md          # Comprehensive audit guide
├── docs/                        # Documentation
│   ├── project_structure.md
│   ├── algorithm.md
│   └── architecture.md
├── .gitignore
├── go.mod
├── LICENSE
├── Makefile
└── README.md
```

## How to Run

### Quick Start

```bash
# Clone and navigate to project
cd tetris-optimizer

# Run with example file
go run ./cmd/tetris-optimizer ./audit/examples/good_example_00.txt
```

### Using Make

```bash
# Build binary
make build

# Run with built binary
./tetris-optimizer ./audit/examples/good_example_00.txt

# Run tests
make test

# Generate coverage report
make coverage
```

### Direct Execution

```bash
# Run directly
go run ./cmd/tetris-optimizer <input_file>

# Build and run
go build -o tetris-optimizer ./cmd/tetris-optimizer
./tetris-optimizer <input_file>
```

## Requirements

- **Go**: Version 1.23 or higher
- **OS**: Linux, macOS, or Windows
- **Dependencies**: None (standard library only)

## Algorithm

The program uses an **optimized backtracking algorithm** with the following steps:

### 1. Parse Input
- Read file and split into 4×4 blocks
- Validate each block (4 connected '#' characters)
- Collect all errors (not just first)
- Normalize coordinates to (0,0) origin
- Assign letter IDs (A, B, C, ...)
- Maximum 26 tetrominoes (A-Z)

### 2. Apply Heuristic
- Sort tetrominoes by constraint (fewer rotations first)
- Most constrained pieces placed first
- Improves backtracking performance

### 3. Calculate Minimum Board Size
```
minimumSize = ceil(sqrt(numberOfTetrominoes × 4))
```
Where numberOfTetrominoes is the count of tetrominoes.

### 4. Generate Rotations
- For each tetromino, generate all unique 90° rotations
- Square pieces: 1 rotation
- Line pieces: 2 rotations
- L/T pieces: 4 rotations

### 5. Backtracking Solver
```
For size = minimumSize to MaxBoardSize (20):
  Create board of current size
  Try to place all tetrominoes using backtracking
  If successful, return board
Return ERROR if no solution found
```

### 6. Optimized Placement
- Find first empty cell on board
- Try all rotations of current tetromino
- For each rotation, try positions from -3 to +3 relative to empty cell
- Recursively place remaining tetrominoes
- Backtrack if placement fails

**Key Optimizations:** 
- First-empty-cell strategy ensures no gaps are left unfilled
- Constraint-based heuristic reduces backtracking depth
- Two-phase validation prevents partial placements
- Rotation precomputation avoids redundant calculations

For detailed algorithm explanation, see [algorithm](docs/algorithm.md).

## Input Format

Each tetromino is defined in a 4×4 grid:

```
....
.##.
.##.
....

.#..
.##.
.#..
....
```

**Rules:**
- Each tetromino: 4 lines × 4 characters
- Use `#` for blocks, `.` for empty spaces
- Exactly 4 `#` characters per tetromino
- Blocks must be connected (up/down/left/right)
- Separate tetrominoes with blank line
- No trailing blank line required

**Valid Tetrominoes:**
- Square: `##` / `##`
- Line: `####` (horizontal or vertical)
- L-shape: `#` / `#` / `##`
- T-shape: `###` / `.#.`
- Z-shape: `.##` / `##.`
- And all rotations/variations

## Output Format

The program outputs the smallest square board with all tetrominoes placed:

```
AA
AA
```

**Format:**
- Square board (N×N)
- Each tetromino labeled with unique letter (A-Z)
- Empty cells marked with `.`
- Each row on separate line
- Trailing newline included

**Example with 2 tetrominoes:**
```
Input:
  A: Square (2×2)
  B: Line (1×4)

Output:
AAB.
AAB.
..B.
..B.
```

**Error Output:**
If input is invalid or no solution exists:
```
ERROR
```

## Testing

### Run All Tests

```bash
# Run unit tests
make test

# Or use go directly
go test -v ./testfiles/...
```

### Test Coverage

```bash
# Generate coverage report
make coverage

# View in browser
go tool cover -html=coverage.out
```

### Audit Tests

```bash
# Test all bad examples 
go run ./cmd/tetris-optimizer ./audit/examples/bad_example_00.txt
go run ./cmd/tetris-optimizer ./audit/examples/bad_example_01.txt
go run ./cmd/tetris-optimizer ./audit/examples/bad_example_02.txt
go run ./cmd/tetris-optimizer ./audit/examples/bad_example_03.txt
go run ./cmd/tetris-optimizer ./audit/examples/bad_example_04.txt
go run ./cmd/tetris-optimizer ./audit/examples/bad_format.txt

# Test all good examples
go run ./cmd/tetris-optimizer ./audit/examples/good_example_00.txt
go run ./cmd/tetris-optimizer ./audit/examples/good_example_01.txt
go run ./cmd/tetris-optimizer ./audit/examples/good_example_02.txt
go run ./cmd/tetris-optimizer ./audit/examples/good_example_03.txt
go run ./cmd/tetris-optimizer ./audit/examples/hard_example.txt
```

For complete audit guide, see [audit guide](audit/audit_guide.md).

## Documentation

- [Project Structure](docs/project_structure.md) - File descriptions and repository layout
- [Architecture](docs/architecture.md) - System design and execution flows
- [Algorithm](docs/algorithm.md) - Algorithm explanation with complexity analysis

## Author

**Katerina Kasdanastasi**

## License

This project is licensed under the [MIT License](LICENSE).
