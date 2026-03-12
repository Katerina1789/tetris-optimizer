# Architecture

## Repository Structure

```
tetris-optimizer/
├── cmd/
│   └── tetris-optimizer/
│       └── main.go                      # Application entry point
├── internal/
│   ├── board/
│   │   ├── board.go                     # Board structure and operations
│   │   └── place.go                     # Placement and removal logic
│   ├── parser/
│   │   ├── parser.go                    # File parsing and tetromino creation
│   │   └── validate.go                  # Input validation logic
│   ├── solver/
│   │   ├── solver.go                    # Main solving algorithm
│   │   └── backtracking.go              # Backtracking implementation
│   └── tetromino/
│       ├── tetromino.go                 # Tetromino structure and creation
│       ├── rotations.go                 # Rotation generation
│       └── normalize.go                 # Coordinate normalization
├── testfiles/
│   ├── board/
│   │   ├── board_test.go                # Tests for board operations
│   │   └── place_test.go                # Tests for placement logic
│   ├── parser/
│   │   ├── parser_test.go               # Tests for file parsing
│   │   └── validate_test.go             # Tests for validation
│   ├── solver/
│   │   ├── solver_test.go               # Tests for solver
│   │   └── backtracking_test.go         # Tests for backtracking
│   └── tetromino/
│       ├── tetromino_test.go            # Tests for tetromino creation
│       ├── rotations_test.go            # Tests for rotations
│       └── normalize_test.go            # Tests for normalization
├── audit/
│   ├── examples/
│   │   ├── bad_example_00.txt           # Invalid input test cases
│   │   ├── bad_example_01.txt
│   │   ├── bad_example_02.txt
│   │   ├── bad_example_03.txt
│   │   ├── bad_example_04.txt
│   │   ├── bad_format.txt
│   │   ├── good_example_00.txt          # Valid input test cases
│   │   ├── good_example_01.txt
│   │   ├── good_example_02.txt
│   │   ├── good_example_03.txt
│   │   └── hard_example.txt             # Performance test case
│   └── audit_guide.md                   # Comprehensive audit guide
├── docs/
│   ├── architecture.md                  # This file - system architecture
│   ├── algorithm.md                     # Algorithm explanation
│   └── flowchart.md                     # Process flow diagrams
├── .gitignore                           # Git exclusion rules
├── go.mod                               # Go module definition
├── LICENSE                              # MIT License
├── Makefile                             # Build automation
└── README.md                            # Project documentation
```

---

## File Descriptions

### **cmd/tetris-optimizer/main.go**
Application entry point that orchestrates the entire solving process.

**Functions:**
- `main()` - Validates command-line arguments, parses input file, runs solver, and prints result or ERROR

**Process Flow:**
1. Check for exactly 2 arguments (program name + file path)
2. Parse tetrominoes from input file using `parser.ParseFile()`
3. Solve using `solver.Solve()`
4. Print board result or ERROR on failure

**Error Handling:**
- Wrong number of arguments → ERROR
- Parse failure → ERROR
- Solve failure → ERROR

---

### **internal/board/board.go**
Defines the square board structure and basic operations.

**Type:**
- `Board` - Struct containing `Size` (int) and `Cells` ([][]rune)

**Functions:**
- `New(size)` - Creates a Size × Size board filled with '.'
- `String()` - Returns board as formatted string with newlines
- `InBounds(r, c)` - Checks if coordinates are within board boundaries
- `IsEmpty(r, c)` - Returns true if cell contains '.'

**Board Representation:**
- Square grid of runes
- Empty cells marked with '.'
- Placed tetrominoes marked with letters A-Z

---

### **internal/board/place.go**
Handles tetromino placement and removal operations.

**Functions:**
- `Place(t, baseR, baseC)` - Attempts to place tetromino at position; returns true if successful
- `Remove(t, baseR, baseC)` - Removes tetromino from board at position
- `FirstEmpty()` - Finds first empty cell (top-left to bottom-right); returns (row, col, found)

**Placement Logic:**
1. Check all blocks fit within bounds
2. Check all target cells are empty
3. If valid, mark cells with tetromino ID
4. Return success/failure

**Optimization:**
- `FirstEmpty()` enables efficient backtracking by finding next placement target

---

### **internal/parser/parser.go**
Reads and parses input files into tetromino structures.

**Functions:**
- `ParseFile(path)` - Reads file and returns slice of tetrominoes labeled A, B, C, ...
- `parseTetrominoBlock(lines, id)` - Converts 4-line block into validated tetromino

**Input Format:**
- Each tetromino: 4 lines of 4 characters
- Characters: '#' (block) or '.' (empty)
- Tetrominoes separated by blank lines
- Must have exactly 4 '#' characters
- Blocks must be connected

**Error Conditions:**
- File not found
- Invalid block dimensions
- Wrong character count
- Disconnected blocks

---

### **internal/parser/validate.go**
Validates tetromino structure and connectivity.

**Type:**
- `point` - Internal struct with `r, c` (int) for BFS traversal

**Functions:**
- `ValidateGrid(grid)` - Ensures grid has exactly 4 connected '#' cells

**Validation Process:**
1. Count '#' characters (must be exactly 4)
2. Perform BFS from first '#' to check connectivity
3. Verify all 4 blocks are reachable
4. Return error if invalid

**Connectivity Check:**
- Uses breadth-first search (BFS)
- Checks 4-directional adjacency (up, down, left, right)
- Ensures single connected component

---

### **internal/solver/solver.go**
Main solving algorithm that finds smallest board size.

**Functions:**
- `Solve(tetros)` - Tries increasing board sizes until solution found

**Algorithm:**
1. Calculate minimum board size: `ceil(sqrt(n × 4))`
2. Try each size from minimum to 20
3. Create board of current size
4. Attempt backtracking placement
5. Return first successful board
6. Error if no solution found

**Size Calculation:**
- n tetrominoes × 4 blocks = total blocks
- Minimum area = total blocks
- Minimum side = sqrt(area) rounded up

---

### **internal/solver/backtracking.go**
Implements optimized backtracking algorithm for tetromino placement.

**Functions:**
- `backtrack(b, tetros, idx)` - Recursively places tetrominoes; returns true if all placed

**Optimization Strategy:**
1. Find first empty cell on board
2. Try all rotations of current tetromino
3. For each rotation, try positions where a block covers the empty cell
4. Recursively place remaining tetrominoes
5. Backtrack if placement fails

**Key Optimization:**
- Only tries positions that cover the first empty cell
- Reduces search space from O(n² × rotations) to O(4 × rotations) per piece
- Ensures no gaps are left unfilled
- Critical for performance on hard examples

**Backtracking Process:**
```
backtrack(board, pieces, index):
  if index == len(pieces): return true
  find first empty cell (r, c)
  for each rotation of pieces[index]:
    for each block in rotation:
      try placing rotation so this block covers (r, c)
      if placement succeeds:
        if backtrack(board, pieces, index+1): return true
        remove placement
  return false
```

---

### **internal/tetromino/tetromino.go**
Defines tetromino structure and creation from grid.

**Types:**
- `Point` - Struct with `R, C` (int) representing block coordinates
- `Tetromino` - Struct with `Blocks` ([]Point) and `ID` (rune)

**Functions:**
- `FromGrid(grid, id)` - Extracts '#' positions from 4×4 grid, normalizes, and creates tetromino

**Process:**
1. Scan grid for '#' characters
2. Collect coordinates as Points
3. Normalize to (0,0) origin
4. Return Tetromino with ID

---

### **internal/tetromino/rotations.go**
Generates all unique 90-degree rotations of a tetromino.

**Functions:**
- `Rotations(t)` - Returns slice of unique rotations (1-4 rotations)
- `rotate90(blocks)` - Rotates points 90° clockwise: (r,c) → (c,-r)
- `keyOf(blocks)` - Creates string key for duplicate detection

**Rotation Logic:**
1. Start with original orientation
2. Rotate 90° three times
3. Normalize each rotation
4. Use string key to detect duplicates
5. Return only unique rotations

**Duplicate Detection:**
- Converts normalized points to byte string
- Uses map to track seen rotations
- Square pieces: 1 unique rotation
- Line pieces: 2 unique rotations
- L/T pieces: 4 unique rotations

---

### **internal/tetromino/normalize.go**
Normalizes tetromino coordinates to top-left origin.

**Functions:**
- `normalize(blocks)` - Shifts all points so minimum row and column are 0

**Normalization Process:**
1. Find minimum row and column values
2. Subtract minimums from all points
3. Return shifted coordinates

**Purpose:**
- Enables rotation comparison
- Simplifies placement logic
- Ensures consistent representation

---

### **testfiles/board/board_test.go**
Unit tests for board structure and basic operations.

**Tests:**
- `TestNew` - Verifies board creation with correct size and empty cells
- `TestString` - Tests string representation formatting
- `TestInBounds` - Validates boundary checking for various coordinates
- `TestIsEmpty` - Checks empty cell detection

---

### **testfiles/board/place_test.go**
Unit tests for placement and removal operations.

**Tests:**
- `TestPlace` - Verifies successful tetromino placement
- `TestPlaceOutOfBounds` - Tests rejection of out-of-bounds placement
- `TestPlaceOverlap` - Tests rejection of overlapping placement
- `TestRemove` - Verifies tetromino removal restores empty cells

---

### **testfiles/parser/parser_test.go**
Unit tests for file parsing functionality.

**Tests:**
- `TestParseFileValid` - Tests successful parsing of valid input
- `TestParseFileInvalidHeight` - Tests error on wrong line count
- `TestParseFileInvalidWidth` - Tests error on wrong line width
- `TestParseFileNonExistent` - Tests error on missing file

---

### **testfiles/parser/validate_test.go**
Unit tests for validation logic.

**Tests:**
- `TestValidateGridValid` - Tests acceptance of valid 4-block connected grid
- `TestValidateGridTooFewBlocks` - Tests rejection of < 4 blocks
- `TestValidateGridTooManyBlocks` - Tests rejection of > 4 blocks
- `TestValidateGridNotConnected` - Tests rejection of disconnected blocks

---

### **testfiles/solver/solver_test.go**
Unit tests for main solving algorithm.

**Tests:**
- `TestSolveSingleTetromino` - Tests solving with one piece
- `TestSolveMultipleTetrominos` - Tests solving with multiple pieces

---

### **testfiles/solver/backtracking_test.go**
Unit tests for backtracking algorithm.

**Tests:**
- `TestBacktrackingSimple` - Tests simple backtracking case
- `TestBacktrackingComplex` - Tests complex multi-piece backtracking

---

### **testfiles/tetromino/tetromino_test.go**
Unit tests for tetromino creation.

**Tests:**
- `TestFromGrid` - Tests tetromino creation from grid with normalization
- `TestFromGridLine` - Tests line-shaped tetromino creation

---

### **testfiles/tetromino/rotations_test.go**
Unit tests for rotation generation.

**Tests:**
- `TestRotationsSquare` - Tests square piece rotations (1-4 unique)
- `TestRotationsLine` - Tests line piece rotations (2-4 unique)
- `TestRotationsL` - Tests L-shaped piece rotations (4 unique)

---

### **testfiles/tetromino/normalize_test.go**
Unit tests for coordinate normalization.

**Tests:**
- `TestNormalizeAlreadyNormalized` - Tests already-normalized coordinates
- `TestNormalizeOffset` - Tests normalization of offset coordinates

---

### **audit/audit_guide.md**
Comprehensive audit guide with 22 test cases covering all requirements.

**Test Categories:**
- Functional Tests (16 tests): Package requirements, error handling, output validation
- Basic Tests (4 tests): Performance, test coverage, code quality
- Social Tests (2 tests): Learning value, recommendation

**Key Tests:**
- All bad examples print ERROR
- All good examples produce correct output
- Performance under 30 seconds for all cases
- Proper tetromino representation

---

### **.gitignore**
Git version control exclusion rules.

**Excluded Items:**
- Binaries (*.exe, *.dll, *.so, *.dylib, tetris-optimizer)
- Test binaries (*.test)
- Coverage reports (*.out)
- IDE files (.vscode/, .idea/, *.swp)
- OS files (.DS_Store, Thumbs.db)

**Preserved Items:**
- All .txt files in audit/examples/

---

### **go.mod**
Go module definition.

**Contents:**
- Module name: `tetris-optimizer`
- Go version: 1.23 or compatible
- No external dependencies (standard library only)

---

### **LICENSE**
MIT License for open-source distribution.

**Details:**
- Copyright 2026 Katerina Kasdanastasi
- Permits commercial and private use
- Requires attribution
- No warranty or liability

---

### **Makefile**
Build automation and testing workflow.

**Targets:**
- `build` - Compile binary to `tetris-optimizer`
- `run` - Execute with go run
- `test` - Run all unit tests with verbose output
- `coverage` - Generate and display test coverage report
- `clean` - Remove generated files
- `all` - Run complete test suite

---

### **README.md**
Comprehensive project documentation and usage guide.

**Contents:**
- Project description and features
- Installation and usage instructions
- Algorithm explanation
- Input format specification
- Testing information
- Author and license details
