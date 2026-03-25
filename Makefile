.PHONY: all build run test coverage clean bench profile

# Build binary
build:
	go build -o tetris-optimizer ./cmd/tetris-optimizer

# Run application
run:
	go run ./cmd/tetris-optimizer

# Run all tests
test:
	go test -v ./testfiles/...

# Run tests with coverage
coverage:
	echo "mode: set" > coverage.out

	go test ./testfiles/board -coverpkg=./internal/board -coverprofile=coverage.tmp
	grep -v "mode:" coverage.tmp >> coverage.out

	go test ./testfiles/parser -coverpkg=./internal/parser -coverprofile=coverage.tmp
	grep -v "mode:" coverage.tmp >> coverage.out

	go test ./testfiles/solver -coverpkg=./internal/solver -coverprofile=coverage.tmp
	grep -v "mode:" coverage.tmp >> coverage.out

	go test ./testfiles/tetromino -coverpkg=./internal/tetromino -coverprofile=coverage.tmp
	grep -v "mode:" coverage.tmp >> coverage.out

	go tool cover -func=coverage.out

# Run benchmarks with CPU and memory stats
bench:
	go test -bench=. -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof ./testfiles/...

# Run tests with CPU and memory profiling
profile:
	go test -v -cpuprofile=cpu.prof -memprofile=mem.prof ./testfiles/...
	@echo "\nTo analyze CPU profile: go tool pprof cpu.prof"
	@echo "To analyze memory profile: go tool pprof mem.prof"

# Clean generated files
clean:
	rm -f tetris-optimizer coverage.out coverage.tmp
	rm -f cpu.prof mem.prof

# Run all checks
all: test
