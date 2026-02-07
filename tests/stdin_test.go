package tests

import (
	"bytes"
	"os"
	"testing"
)

// TestReadFromStdin verifies that the program correctly reads
// a single integer from standard input.
func TestReadFromStdin(t *testing.T) {
	// Arrange
	input := "189\n"
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	_, _ = w.Write([]byte(input))
	_ = w.Close()
	os.Stdin = r

	// Act
	value, err := readInput()

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if value != 189 {
		t.Fatalf("expected 189, got %d", value)
	}
}

// TestReadFromStdinInvalidInput ensures invalid input is handled safely.
func TestReadFromStdinInvalidInput(t *testing.T) {
	// Arrange
	input := "abc\n"
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	_, _ = w.Write([]byte(input))
	_ = w.Close()
	os.Stdin = r

	// Act
	_, err = readInput()

	// Assert
	if err == nil {
		t.Fatalf("expected error for invalid input, got nil")
	}
}

// BenchmarkReadFromStdin measures stdin reading performance.
func BenchmarkReadFromStdin(b *testing.B) {
	input := "123\n"

	for i := 0; i < b.N; i++ {
		r := bytes.NewBufferString(input)
		oldStdin := os.Stdin
		os.Stdin = os.NewFile(uintptr(0), "/dev/stdin")
		os.Stdin = r

		_, _ = readInput()
		os.Stdin = oldStdin
	}
}
