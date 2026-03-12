.PHONY: all build run test coverage clean

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
	go test -coverprofile=coverage.out ./testfiles/...
	go tool cover -func coverage.out

# Clean generated files
clean:
	rm -f tetris-optimizer coverage.out

# Run all checks
all: test
