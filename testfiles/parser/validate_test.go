package parser_test

import (
	"testing"

	"tetris-optimizer/internal/parser"
)

func TestValidateGridValid(t *testing.T) {
	grid := [][]rune{
		[]rune("##.."),
		[]rune("##.."),
		[]rune("...."),
		[]rune("...."),
	}

	if err := parser.ValidateGrid(grid, 1); err != nil {
		t.Fatalf("expected valid grid, got %v", err)
	}
}

func TestValidateGridWrongBlockCount(t *testing.T) {
	grid := [][]rune{
		[]rune("#..."),
		[]rune("#..."),
		[]rune("#..."),
		[]rune("...."),
	}

	if err := parser.ValidateGrid(grid, 1); err == nil {
		t.Fatalf("expected error for wrong block count")
	}
}

func TestValidateGridDisconnected(t *testing.T) {
	grid := [][]rune{
		[]rune("##.."),
		[]rune("...."),
		[]rune("..##"),
		[]rune("...."),
	}

	if err := parser.ValidateGrid(grid, 1); err == nil {
		t.Fatalf("expected error for disconnected blocks")
	}
}
