package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestAsciiMapHasCapitalLetters(t *testing.T) {
	for char := range ascii {
		if char < 'A' || char > 'Z' {
			t.Errorf("expected capital letter, got %q", char)
		}
	}
}

func TestAsciiMapHeight(t *testing.T) {
	shape := ascii['A']
	height := len(shape)

	if height != 3 {
		t.Errorf("expected height 3, got %d", height)
	}
}

func TestAsciiMapWidth(t *testing.T) {
	shape := ascii['A']
	expectedWidth := 6

	for rowIndex, row := range shape {
		width := len(row)
		if width != expectedWidth {
			t.Errorf("row %d: expected width %d, got %d", rowIndex, expectedWidth, width)
		}
	}
}

func TestAsciiMapContainsA(t *testing.T) {
	if _, ok := ascii['A']; !ok {
		t.Fatal("expected ascii map to contain 'A'")
	}
}

func TestAsciiMapDoesNotContainLowercase(t *testing.T) {
	if _, ok := ascii['a']; ok {
		t.Error("did not expect ascii map to contain lowercase 'a'")
	}
}

func TestAsciiRowsAreNotEmpty(t *testing.T) {
	for char, shape := range ascii {
		for rowIndex, row := range shape {
			if row == "" {
				t.Errorf("character %q row %d is empty", char, rowIndex)
			}
		}
	}
}

func TestAsciiRowsHaveSameWidth(t *testing.T) {
	for char, shape := range ascii {
		if len(shape) == 0 {
			t.Errorf("character %q has no rows", char)
			continue
		}

		expectedWidth := len(shape[0])

		for rowIndex, row := range shape {
			if len(row) != expectedWidth {
				t.Errorf("character %q row %d: expected width %d, got %d",
					char, rowIndex, expectedWidth, len(row))
			}
		}
	}
}

func TestPrintShapeMissingCharacter(t *testing.T) {
	shape := ascii['B']

	if len(shape) != 0 {
		t.Errorf("expected missing character 'B' to have empty shape, got %v", shape)
	}
}

func TestPrintShapePrintsHeightWidthAndRows(t *testing.T) {
	output := captureOutput(func() {
		printShap('A')
	})

	expectedParts := []string{
		"height: 3",
		"width: 6",
		"AzzzzA",
		"A    A",
	}

	for _, part := range expectedParts {
		if !strings.Contains(output, part) {
			t.Errorf("expected output to contain %q, got %q", part, output)
		}
	}
}

func captureOutput(run func()) string {
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	run()

	writer.Close()
	os.Stdout = oldStdout

	var buffer bytes.Buffer
	io.Copy(&buffer, reader)
	return buffer.String()
}
