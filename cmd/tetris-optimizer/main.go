package main

// Main entrypoint: parses args, loads tetrominoes, runs solver, prints result or ERROR.

import (
	"fmt"
	"os"

	"tetris-optimizer/internal/parser"
	"tetris-optimizer/internal/solver"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Input should be: go run ./cmd/tetris-optimizer <input_file>")
		return
	}

	path := os.Args[1]
	tetros, err := parser.ParseFile(path)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	if len(tetros) == 0 {
		fmt.Println("ERROR: No tetrominoes found in file")
		return
	}

	board, err := solver.Solve(tetros)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	fmt.Print(board.String())
}
