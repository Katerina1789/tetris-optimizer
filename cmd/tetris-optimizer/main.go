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
		fmt.Println("ERROR")
		return
	}

	path := os.Args[1]
	tetros, err := parser.ParseFile(path)
	if err != nil || len(tetros) == 0 {
		fmt.Println("ERROR")
		return
	}

	board, err := solver.Solve(tetros)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	fmt.Print(board.String())
}
