package parser_test

import (
	"os"
	"testing"

	"tetris-optimizer/internal/parser"
)

func writeTemp(t *testing.T, content string) string {
	t.Helper()
	tempFile, err := os.CreateTemp("", "tetro*.txt")
	if err != nil {
		t.Fatalf("temp file error: %v", err)
	}
	if _, err := tempFile.WriteString(content); err != nil {
		t.Fatalf("write error: %v", err)
	}
	tempFile.Close()
	return tempFile.Name()
}

func TestParseFileSingleValid(t *testing.T) {
	content := `##..
##..
....
....
`
	filePath := writeTemp(t, content)
	defer os.Remove(filePath)

	tetrominoes, err := parser.ParseFile(filePath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tetrominoes) != 1 {
		t.Fatalf("expected 1 tetromino, got %d", len(tetrominoes))
	}
	if tetrominoes[0].ID != 'A' {
		t.Fatalf("expected ID A, got %c", tetrominoes[0].ID)
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
	filePath := writeTemp(t, content)
	defer os.Remove(filePath)

	tetrominoes, err := parser.ParseFile(filePath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tetrominoes) != 2 {
		t.Fatalf("expected 2 tetrominoes, got %d", len(tetrominoes))
	}
}

func TestParseFileInvalidHeight(t *testing.T) {
	content := `##..
##..
....
` // only 3 lines

	filePath := writeTemp(t, content)
	defer os.Remove(filePath)

	if _, err := parser.ParseFile(filePath); err == nil {
		t.Fatalf("expected height error")
	}
}

func TestParseFileInvalidWidth(t *testing.T) {
	content := `###.
##..
.....
....
`
	filePath := writeTemp(t, content)
	defer os.Remove(filePath)

	if _, err := parser.ParseFile(filePath); err == nil {
		t.Fatalf("expected width error")
	}
}

func TestParseFileInvalidChar(t *testing.T) {
	content := `#x..
##..
....
....
`
	filePath := writeTemp(t, content)
	defer os.Remove(filePath)

	if _, err := parser.ParseFile(filePath); err == nil {
		t.Fatalf("expected invalid char error")
	}
}
