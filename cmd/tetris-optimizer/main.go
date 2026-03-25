package main

// Main entrypoint: parses arguments, loads tetrominoes, runs solver, prints result or ERROR.

import (
	"fmt"
	"os"

	"tetris-optimizer/internal/parser"
	"tetris-optimizer/internal/solver"
)

func main() {
	// Validate command-line arguments
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Usage: tetris-optimizer <input_file>")
		os.Exit(1)
	}

	// Parse tetrominoes from input file
	filePath := os.Args[1]
	tetrominoes, err := parser.ParseFile(filePath)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	// Solve the tetromino packing problem
	resultBoard, err := solver.Solve(tetrominoes)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	// Print the solution
	fmt.Print(resultBoard.String())
}
