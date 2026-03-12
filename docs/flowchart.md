# Flowchart

## Program Execution Flow

```
main()
  ├─ Validate arguments (argc == 2)
  │   └─ NO → Print "ERROR" and exit
  ├─ Extract file path from args[1]
  ├─ ParseFile(path)
  │   ├─ Open file
  │   │   └─ ERROR → Return error
  │   ├─ Read lines and split into blocks
  │   ├─ For each block:
  │   │   ├─ Validate 4×4 format
  │   │   ├─ Validate 4 '#' characters
  │   │   ├─ Validate connectivity (BFS)
  │   │   └─ Create Tetromino with ID
  │   └─ Return tetrominoes array
  ├─ Check parse result
  │   └─ ERROR or empty → Print "ERROR" and exit
  ├─ Solve(tetrominoes)
  │   ├─ Calculate minSize
  │   ├─ For size = minSize to 20:
  │   │   ├─ Create board of current size
  │   │   ├─ Backtrack(board, tetros, 0)
  │   │   └─ If success → Return board
  │   └─ No solution → Return error
  ├─ Check solve result
  │   └─ ERROR → Print "ERROR" and exit
  └─ Print board.String()
```

---

## File Parsing Flow

```
ParseFile(path)
  ├─ Open file
  │   └─ ERROR → Return (nil, error)
  ├─ Create scanner
  ├─ Initialize: blocks=[], current=[]
  ├─ For each line:
  │   ├─ If line is blank:
  │   │   ├─ If current not empty:
  │   │   │   ├─ Append current to blocks
  │   │   │   └─ Reset current = []
  │   │   └─ Continue
  │   └─ Append line to current
  ├─ If current not empty:
  │   └─ Append current to blocks
  ├─ For each block (index i):
  │   ├─ parseTetrominoBlock(block, 'A'+i)
  │   │   ├─ Validate len(lines) == 4
  │   │   │   └─ NO → Return error
  │   │   ├─ For each line:
  │   │   │   ├─ Validate len(line) == 4
  │   │   │   │   └─ NO → Return error
  │   │   │   └─ Validate chars ∈ {'#', '.'}
  │   │   │       └─ NO → Return error
  │   │   ├─ ValidateGrid(grid)
  │   │   │   └─ ERROR → Return error
  │   │   └─ Return FromGrid(grid, id)
  │   └─ Append tetromino to results
  └─ Return (tetrominoes, nil)
```

---

## Validation Flow

```
ValidateGrid(grid)
  ├─ Count '#' characters
  │   └─ If count != 4 → Return error
  ├─ Find first '#' position
  ├─ Initialize BFS:
  │   ├─ visited = {}
  │   ├─ queue = [first position]
  │   └─ visited[first] = true
  ├─ While queue not empty:
  │   ├─ current = queue.pop()
  │   ├─ For each direction (up, down, left, right):
  │   │   ├─ neighbor = current + direction
  │   │   ├─ If out of bounds → Continue
  │   │   ├─ If not '#' → Continue
  │   │   ├─ If already visited → Continue
  │   │   ├─ Mark visited
  │   │   └─ Add to queue
  │   └─ Continue
  ├─ Check len(visited) == 4
  │   └─ NO → Return error (not connected)
  └─ Return nil (valid)
```

---

## Tetromino Creation Flow

```
FromGrid(grid, id)
  ├─ Initialize points = []
  ├─ For r = 0 to 3:
  │   └─ For c = 0 to 3:
  │       └─ If grid[r][c] == '#':
  │           └─ Append Point{r, c} to points
  ├─ Normalize(points)
  │   ├─ Find minR, minC
  │   └─ Shift all points: (r-minR, c-minC)
  └─ Return Tetromino{Blocks: points, ID: id}
```

---

## Solver Flow

```
Solve(tetrominoes)
  ├─ n = len(tetrominoes)
  ├─ minSize = ceil(sqrt(n × 4))
  ├─ For size = minSize to 20:
  │   ├─ board = New(size)
  │   ├─ success = Backtrack(board, tetrominoes, 0)
  │   └─ If success:
  │       └─ Return (board, nil)
  └─ Return (nil, error)
```

---

## Backtracking Flow

```
Backtrack(board, tetrominoes, index)
  ├─ Base case: index == len(tetrominoes)
  │   └─ Return true (all placed)
  │
  ├─ Find first empty cell
  │   ├─ (r, c, found) = board.FirstEmpty()
  │   └─ If not found:
  │       └─ Return index == len(tetrominoes)
  │
  ├─ Get current tetromino
  │   └─ tetromino = tetrominoes[index]
  │
  ├─ Generate rotations
  │   └─ rotations = Rotations(tetromino)
  │
  ├─ For each rotation:
  │   ├─ For each block in rotation:
  │   │   ├─ Calculate base position:
  │   │   │   ├─ baseR = r - block.R
  │   │   │   └─ baseC = c - block.C
  │   │   │
  │   │   ├─ Try placement:
  │   │   │   └─ If board.Place(rotation, baseR, baseC):
  │   │   │       ├─ Recurse:
  │   │   │       │   └─ If Backtrack(board, tetros, index+1):
  │   │   │       │       └─ Return true (solution found)
  │   │   │       └─ Backtrack:
  │   │   │           └─ board.Remove(rotation, baseR, baseC)
  │   │   └─ Continue to next position
  │   └─ Continue to next rotation
  │
  └─ Return false (no solution with current piece)
```

---

## Placement Flow

```
board.Place(tetromino, baseR, baseC)
  ├─ Validation phase:
  │   └─ For each block in tetromino:
  │       ├─ Calculate position: (baseR + block.R, baseC + block.C)
  │       ├─ Check InBounds(r, c)
  │       │   └─ NO → Return false
  │       └─ Check IsEmpty(r, c)
  │           └─ NO → Return false
  │
  ├─ Placement phase:
  │   └─ For each block in tetromino:
  │       ├─ Calculate position: (baseR + block.R, baseC + block.C)
  │       └─ Set board.Cells[r][c] = tetromino.ID
  │
  └─ Return true
```

---

## Removal Flow

```
board.Remove(tetromino, baseR, baseC)
  └─ For each block in tetromino:
      ├─ Calculate position: (baseR + block.R, baseC + block.C)
      ├─ If InBounds(r, c) AND Cells[r][c] == tetromino.ID:
      │   └─ Set Cells[r][c] = '.'
      └─ Continue
```

---

## Rotation Generation Flow

```
Rotations(tetromino)
  ├─ Initialize:
  │   ├─ results = []
  │   ├─ seen = {}
  │   └─ current = tetromino.Blocks
  │
  ├─ For i = 0 to 3:
  │   ├─ Normalize current orientation
  │   ├─ Generate unique key
  │   ├─ If key not in seen:
  │   │   ├─ Add key to seen
  │   │   └─ Append Tetromino{Blocks: normalized, ID: id} to results
  │   └─ Rotate current 90° clockwise:
  │       └─ For each point (r, c):
  │           └─ Transform to (c, -r)
  │
  └─ Return results
```

---

## First Empty Cell Flow

```
board.FirstEmpty()
  ├─ For r = 0 to board.Size-1:
  │   └─ For c = 0 to board.Size-1:
  │       └─ If board.Cells[r][c] == '.':
  │           └─ Return (r, c, true)
  │
  └─ Return (0, 0, false)
```

---

## Example: Solving 2 Tetrominoes

```
Input:
  Tetromino A: Square (##/##)
  Tetromino B: Line (####)

Execution:
  ├─ Parse: A=[{0,0},{0,1},{1,0},{1,1}], B=[{0,0},{0,1},{0,2},{0,3}]
  ├─ minSize = ceil(sqrt(8)) = 3
  │
  ├─ Try size 3:
  │   ├─ Board: 3×3 (9 cells, need 8)
  │   ├─ Backtrack(board, [A,B], 0):
  │   │   ├─ FirstEmpty() → (0,0)
  │   │   ├─ Try A at (0,0):
  │   │   │   ├─ Place A: AA./AA./...
  │   │   │   ├─ Backtrack(board, [A,B], 1):
  │   │   │   │   ├─ FirstEmpty() → (0,2)
  │   │   │   │   ├─ Try B vertical: Fails (extends beyond)
  │   │   │   │   ├─ Try B horizontal: Fails (only 1 cell)
  │   │   │   │   └─ Return false
  │   │   │   └─ Remove A
  │   │   └─ Try other positions... (all fail)
  │   └─ Return false
  │
  ├─ Try size 4:
  │   ├─ Board: 4×4 (16 cells, need 8)
  │   ├─ Backtrack(board, [A,B], 0):
  │   │   ├─ FirstEmpty() → (0,0)
  │   │   ├─ Try A at (0,0):
  │   │   │   ├─ Place A: AA../AA../..../....
  │   │   │   ├─ Backtrack(board, [A,B], 1):
  │   │   │   │   ├─ FirstEmpty() → (0,2)
  │   │   │   │   ├─ Try B vertical at (0,2):
  │   │   │   │   │   ├─ Place B: AAB./AAB./..B./..B.
  │   │   │   │   │   ├─ Backtrack(board, [A,B], 2):
  │   │   │   │   │   │   └─ index == 2 → Return true
  │   │   │   │   │   └─ Return true
  │   │   │   │   └─ Return true
  │   │   │   └─ Return true
  │   │   └─ Return true
  │   └─ Solution found!
  │
  └─ Output:
      AAB.
      AAB.
      ..B.
      ..B.
```

---

## Performance Optimization Points

### 1. First Empty Cell Strategy
```
Without optimization:
  For each piece:
    Try all N² positions → O(N²) per piece

With optimization:
  For each piece:
    Find first empty cell → O(N²) once
    Try only positions covering that cell → O(4) per piece
```

### 2. Rotation Caching
```
Without caching:
  For each placement attempt:
    Generate rotations → O(4) per attempt

With caching:
  Once per piece:
    Generate rotations → O(4) total
  Reuse for all attempts → O(1) per attempt
```

### 3. Early Termination
```
Without early termination:
  Find all solutions
  Select smallest board

With early termination:
  Iterate sizes from smallest
  Return first solution
  Guaranteed optimal
```

---

## Error Handling Flow

```
Error Detection
  ├─ Invalid arguments
  │   └─ Print "ERROR" and exit
  ├─ File not found
  │   └─ Print "ERROR" and exit
  ├─ Parse error
  │   ├─ Invalid dimensions
  │   ├─ Invalid characters
  │   ├─ Wrong block count
  │   └─ Disconnected blocks
  │       └─ Print "ERROR" and exit
  └─ No solution found
      └─ Print "ERROR" and exit
```

---

## Testing Flow

```
Test Execution
  ├─ Bad Examples:
  │   ├─ bad_example_00.txt → ERROR
  │   ├─ bad_example_01.txt → ERROR
  │   ├─ bad_example_02.txt → ERROR
  │   ├─ bad_example_03.txt → ERROR
  │   ├─ bad_example_04.txt → ERROR
  │   └─ bad_format.txt → ERROR
  │
  ├─ Good Examples:
  │   ├─ good_example_00.txt → Valid board
  │   ├─ good_example_01.txt → Valid board
  │   ├─ good_example_02.txt → Valid board (< 30s)
  │   ├─ good_example_03.txt → Valid board (< 30s)
  │   └─ hard_example.txt → Valid board (< 30s)
  │
  └─ Unit Tests:
      ├─ board/ → 8 tests
      ├─ parser/ → 8 tests
      ├─ solver/ → 4 tests
      └─ tetromino/ → 7 tests
```
