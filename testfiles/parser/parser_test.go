package parser_test

import (
	"os"
	"testing"
	"tetris-optimizer/internal/parser"
)

func writeTemp(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "tetro*.txt")
	if err != nil {
		t.Fatalf("temp file error: %v", err)
	}
	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("write error: %v", err)
	}
	f.Close()
	return f.Name()
}

func TestParseFileSingleValid(t *testing.T) {
	content := `##..
##..
....
....
`
	path := writeTemp(t, content)
	defer os.Remove(path)

	tetros, err := parser.ParseFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tetros) != 1 {
		t.Fatalf("expected 1 tetromino, got %d", len(tetros))
	}
	if tetros[0].ID != 'A' {
		t.Fatalf("expected ID A, got %c", tetros[0].ID)
	}
}

func TestParseFileMultiple(t *testing.T) {
	content := `##..
##..
....
....

#...
#...
#...
#...
`
	path := writeTemp(t, content)
	defer os.Remove(path)

	tetros, err := parser.ParseFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tetros) != 2 {
		t.Fatalf("expected 2 tetrominoes, got %d", len(tetros))
	}
}

func TestParseFileInvalidHeight(t *testing.T) {
	content := `##..
##..
....
` // only 3 lines

	path := writeTemp(t, content)
	defer os.Remove(path)

	if _, err := parser.ParseFile(path); err == nil {
		t.Fatalf("expected height error")
	}
}

func TestParseFileInvalidWidth(t *testing.T) {
	content := `###.
##..
.....
....
`
	path := writeTemp(t, content)
	defer os.Remove(path)

	if _, err := parser.ParseFile(path); err == nil {
		t.Fatalf("expected width error")
	}
}

func TestParseFileInvalidChar(t *testing.T) {
	content := `#x..
##..
....
....
`
	path := writeTemp(t, content)
	defer os.Remove(path)

	if _, err := parser.ParseFile(path); err == nil {
		t.Fatalf("expected invalid char error")
	}
}
