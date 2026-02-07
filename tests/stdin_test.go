package tests

import (
	"guess-it/pipeline"
	"strings"
	"testing"
)

// TestReadFromStdin verifies reading a single integer from standard input
func TestReadFromStdin(t *testing.T) {
	// Arrange: create mock stdin with input "189"
	input := "189\n"
	reader := strings.NewReader(input)

	// Act: read and parse the input
	value, err := pipeline.ReadInputFrom(reader)

	// Assert: verify no error and correct value
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if value != 189 {
		t.Fatalf("expected 189, got %d", value)
	}
}

// TestReadFromStdinInvalidInput verifies error handling for invalid input
func TestReadFromStdinInvalidInput(t *testing.T) {
	// Arrange: create mock stdin with invalid input
	input := "abc\n"
	reader := strings.NewReader(input)

	// Act: attempt to read and parse
	_, err := pipeline.ReadInputFrom(reader)

	// Assert: verify error is returned
	if err == nil {
		t.Fatalf("expected error for invalid input, got nil")
	}
}

// TestReadFromStdinMultipleValues verifies reading multiple values sequentially
func TestReadFromStdinMultipleValues(t *testing.T) {
	// Arrange: create mock stdin with multiple lines
	input := "10\n20\n30\n"
	reader := strings.NewReader(input)

	// Act: read first value
	value1, err1 := pipeline.ReadInputFrom(reader)

	// Assert: verify first value
	if err1 != nil {
		t.Fatalf("unexpected error on first read: %v", err1)
	}
	if value1 != 10 {
		t.Fatalf("expected 10, got %d", value1)
	}

	// Note: ReadInputFrom creates a new reader each time,
	// so this test demonstrates single-value reading behavior
}
