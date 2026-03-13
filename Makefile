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

# Clean generated files
clean:
	rm -f tetris-optimizer coverage.out
	rm -f coverage.tmp

# Run all checks
all: test
