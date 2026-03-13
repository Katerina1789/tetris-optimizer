# Audit Guide: Tetris-Optimizer

## Functional Tests

### Test 1: Package Requirements
**Check:** Has the requirement for allowed packages been respected?

**How to Test:**
```bash
# Check go.mod for dependencies
cat go.mod
# Expected: Only standard Go packages (no external dependencies)
```

**Checklist:**
- [ ] Only standard library packages used
- [ ] No third-party dependencies in go.mod

---

### Test 2: Bad Example 00
**Check:** Does the program print ERROR?

**How to Test:**
```bash
go run ./cmd/tetris-optimizer ./audit/examples/bad_example_00.txt
# Expected: ERROR
```

**Checklist:**
- [ ] Prints "ERROR"
- [ ] No crash or panic

---

### Test 3: Bad Example 01
**Check:** Does the program print ERROR?

**How to Test:**
```bash
go run ./cmd/tetris-optimizer ./audit/examples/bad_example_01.txt
# Expected: ERROR
```

**Checklist:**
- [ ] Prints "ERROR"
- [ ] No crash or panic

---

### Test 4: Bad Example 02
**Check:** Does the program print ERROR?

**How to Test:**
```bash
go run ./cmd/tetris-optimizer ./audit/examples/bad_example_02.txt
# Expected: ERROR
```

**Checklist:**
- [ ] Prints "ERROR"
- [ ] No crash or panic

---

### Test 5: Bad Example 03
**Check:** Does the program print ERROR?

**How to Test:**
```bash
go run ./cmd/tetris-optimizer ./audit/examples/bad_example_03.txt
# Expected: ERROR
```

**Checklist:**
- [ ] Prints "ERROR"
- [ ] No crash or panic

---

### Test 6: Bad Example 04
**Check:** Does the program print ERROR?

**How to Test:**
```bash
go run ./cmd/tetris-optimizer ./audit/examples/bad_example_04.txt
# Expected: ERROR
```

**Checklist:**
- [ ] Prints "ERROR"
- [ ] No crash or panic

---

### Test 7: Bad Format
**Check:** Does the program print ERROR?

**How to Test:**
```bash
go run ./cmd/tetris-optimizer ./audit/examples/bad_format.txt
# Expected: ERROR
```

**Checklist:**
- [ ] Prints "ERROR"
- [ ] No crash or panic

---

### Test 8: Good Example 00
**Check:** Does the result contain 0 empty spaces (0 '.')?

**How to Test:**
```bash
go run ./cmd/tetris-optimizer ./audit/examples/good_example_00.txt
# Count dots in output
go run ./cmd/tetris-optimizer ./audit/examples/good_example_00.txt | grep -o '\.' | wc -l
# Expected: 0
```

**Checklist:**
- [ ] Output contains 0 dots
- [ ] All tetrominoes placed
- [ ] Valid square board

---

### Test 9: Good Example 01
**Check:** Does the result contain 9 empty spaces (9 '.')?

**How to Test:**
```bash
go run ./cmd/tetris-optimizer ./audit/examples/good_example_01.txt
# Count dots in output
go run ./cmd/tetris-optimizer ./audit/examples/good_example_01.txt | grep -o '\.' | wc -l
# Expected: 9
```

**Checklist:**
- [ ] Output contains 9 dots
- [ ] All tetrominoes placed
- [ ] Valid square board

---

### Test 10: Good Example 02
**Check:** Does the result contain 4 empty spaces (4 '.') and time limit ≤ 30 seconds?

**How to Test:**
```bash
time go run ./cmd/tetris-optimizer ./audit/examples/good_example_02.txt
# Count dots in output
go run ./cmd/tetris-optimizer ./audit/examples/good_example_02.txt | grep -o '\.' | wc -l
# Expected: 4 dots, time < 30s
```

**Checklist:**
- [ ] Output contains 4 dots
- [ ] Execution time ≤ 30 seconds
- [ ] All tetrominoes placed
- [ ] Valid square board

---

### Test 11: Good Example 03
**Check:** Does the result contain 5 empty spaces (5 '.') and time limit ≤ 30 seconds?

**How to Test:**
```bash
time go run ./cmd/tetris-optimizer ./audit/examples/good_example_03.txt
# Count dots in output
go run ./cmd/tetris-optimizer ./audit/examples/good_example_03.txt | grep -o '\.' | wc -l
# Expected: 5 dots, time < 30s
```

**Checklist:**
- [ ] Output contains 5 dots
- [ ] Execution time ≤ 30 seconds
- [ ] All tetrominoes placed
- [ ] Valid square board

---

### Test 12: Hard Example
**Check:** Does the result contain 1 empty space (1 '.') and time limit ≤ 30 seconds?

**How to Test:**
```bash
time go run ./cmd/tetris-optimizer ./audit/examples/hard_example.txt
# Count dots in output
go run ./cmd/tetris-optimizer ./audit/examples/hard_example.txt | grep -o '\.' | wc -l
# Expected: 1 dot, time < 30s
```

**Checklist:**
- [ ] Output contains 1 dot
- [ ] Execution time ≤ 30 seconds
- [ ] All tetrominoes placed
- [ ] Valid square board

---

### Test 13: Tetromino Presence
**Check:** Are all tetrominoes from the input file present in the output?

**How to Test:**
```bash
# Count tetrominoes in input
grep -c '^####' ./audit/examples/good_example_01.txt

# Count unique letters in output
go run ./cmd/tetris-optimizer ./audit/examples/good_example_01.txt | grep -o '[A-Z]' | sort -u | wc -l
```

**Checklist:**
- [ ] All input tetrominoes appear in output
- [ ] No missing pieces
- [ ] Correct count

---

### Test 14: Character Uniqueness
**Check:** Do different characters correspond to different tetrominoes?

**How to Test:**
```bash
# Visual inspection of output
go run ./cmd/tetris-optimizer ./audit/examples/good_example_01.txt
```

**Checklist:**
- [ ] Each tetromino has unique character (A, B, C, ...)
- [ ] No character reuse
- [ ] Clear visual distinction

---

### Test 15: Single Character Per Tetromino
**Check:** Does one tetromino have only one character?

**How to Test:**
```bash
# Visual inspection
go run ./cmd/tetris-optimizer ./audit/examples/good_example_00.txt
```

**Checklist:**
- [ ] Each tetromino uses exactly one letter
- [ ] No mixed characters in single piece
- [ ] Consistent labeling

---

### Test 16: Project Standards
**Check:** Is this project up to every standard?

**Evaluation Criteria:**
- [ ] **Empty Work:** Project is complete and functional
- [ ] **Incomplete Work:** All features implemented
- [ ] **Invalid Compilation:** Builds successfully with `go run`
- [ ] **Cheating:** Original work
- [ ] **Crashing:** Runs without crashes
- [ ] **Leaks:** No resource leaks

**How to Test:**
```bash
# Test compilation
go run ./cmd/tetris-optimizer ./audit/examples/good_example_00.txt

# Test all examples
for f in ./audit/examples/*.txt; do
  echo "Testing $f"
  go run ./cmd/tetris-optimizer "$f"
done
```

---

## Basic Tests

### Test 17: Performance
**Check:** Does the project run quickly and effectively?

**How to Test:**
```bash
# Test response time for all examples
time go run ./cmd/tetris-optimizer ./audit/examples/good_example_00.txt
time go run ./cmd/tetris-optimizer ./audit/examples/good_example_01.txt
time go run ./cmd/tetris-optimizer ./audit/examples/good_example_02.txt
time go run ./cmd/tetris-optimizer ./audit/examples/good_example_03.txt
time go run ./cmd/tetris-optimizer ./audit/examples/hard_example.txt

```

**Checklist:**
- [ ] Fast execution times
- [ ] No unnecessary data requests
- [ ] Efficient backtracking algorithm
- [ ] Proper recursion usage
- [ ] No performance bottlenecks

---

### Test 18: Test Files
**Check:** Is there a test file for this code?

**How to Test:**
```bash
# Check for test files
ls testfiles/

# Run tests
go test ./testfiles/...
```

**Checklist:**
- [ ] Test files exist in `testfiles/`
- [ ] Tests for board logic
- [ ] Tests for parser logic
- [ ] Tests for solver logic
- [ ] Tests for tetromino logic
- [ ] All tests pass

---

### Test 19: Test Coverage
**Check:** Are the tests checking each possible case?

**How to Test:**
```bash
# Run tests with coverage
go test -cover ./testfiles/...
```

**Test Cases to Check:**
- [ ] Valid tetromino parsing
- [ ] Invalid input handling
- [ ] Board placement logic
- [ ] Rotation generation
- [ ] Normalization
- [ ] Backtracking algorithm
- [ ] Edge cases covered

---

### Test 20: Code Quality
**Check:** Does the code obey good practices?

**How to Test:**
```bash
# Review code structure
ls -R

# Check for organization
cat internal/solver/solver.go
cat internal/parser/parser.go
```

**Code Review Checklist:**
- [ ] Clear package structure
- [ ] Separation of concerns
- [ ] Proper error handling
- [ ] Readable function names
- [ ] No code duplication
- [ ] Comments where needed
- [ ] Follows Go conventions

---

## Social

### Test 21: Learning Value
**Check:** Did you learn anything from this project?

**Discussion Points:**
- Backtracking algorithms
- Tetromino rotation and normalization
- Board placement optimization
- Go package structure
- Efficient recursion

---

### Test 22: Recommendation
**Check:** Would you recommend/nominate this program as an example?

**Evaluation:**
- [ ] Meets all requirements
- [ ] Clean implementation
- [ ] Good documentation
- [ ] Follows best practices
- [ ] Production-ready

---

## Quick Test Commands
```bash
#Test all bad examples
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

# Run tests
go test ./testfiles/...

# Build binary
go build -o tetris-optimizer ./cmd/tetris-optimizer
```

---

## Checklist for Auditors

### Functional
- [ ] Only standard Go packages used
- [ ] bad_example_00.txt prints ERROR
- [ ] bad_example_01.txt prints ERROR
- [ ] bad_example_02.txt prints ERROR
- [ ] bad_example_03.txt prints ERROR
- [ ] bad_example_04.txt prints ERROR
- [ ] bad_format.txt prints ERROR
- [ ] good_example_00.txt: 0 empty spaces
- [ ] good_example_01.txt: 9 empty spaces
- [ ] good_example_02.txt: 4 empty spaces, ≤ 30s
- [ ] good_example_03.txt: 5 empty spaces, ≤ 30s
- [ ] hard_example.txt: 1 empty space, ≤ 30s
- [ ] All tetrominoes present in output
- [ ] Different characters for different tetrominoes
- [ ] One character per tetromino
- [ ] Project meets all standards

### Basic
- [ ] Project runs quickly and effectively
- [ ] Test files exist
- [ ] Tests check all cases
- [ ] Code follows good practices

### Social
- [ ] Project demonstrates learning
- [ ] Worthy of recommendation
