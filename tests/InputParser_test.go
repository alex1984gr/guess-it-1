package tests

import (
	"guess-it/pipeline"
	"testing"
)

// TestInputParserValidInteger verifies parsing of a valid integer string
func TestInputParserValidInteger(t *testing.T) {
	// Create a new parser instance
	parser := pipeline.NewInputParser()

	// Parse a valid integer string
	value, err := parser.Parse("189")

	// Verify no error occurred
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Verify correct value was parsed
	if value != 189 {
		t.Fatalf("expected 189, got %d", value)
	}
}

// TestInputParserWithWhitespace verifies whitespace trimming
func TestInputParserWithWhitespace(t *testing.T) {
	parser := pipeline.NewInputParser()

	// Parse string with leading and trailing spaces
	value, err := parser.Parse("  113  ")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if value != 113 {
		t.Fatalf("expected 113, got %d", value)
	}
}

// TestInputParserInvalidInput verifies error handling for non-numeric input
func TestInputParserInvalidInput(t *testing.T) {
	parser := pipeline.NewInputParser()

	// Attempt to parse invalid input
	_, err := parser.Parse("abc")

	// Verify error is returned
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// TestInputParserEmptyInput verifies error handling for empty input
func TestInputParserEmptyInput(t *testing.T) {
	parser := pipeline.NewInputParser()

	// Attempt to parse empty string
	_, err := parser.Parse("")

	// Verify error is returned
	if err == nil {
		t.Fatalf("expected error for empty input")
	}
}

// TestInputParserNewlineInput verifies handling of newline characters
func TestInputParserNewlineInput(t *testing.T) {
	parser := pipeline.NewInputParser()

	// Parse string with newline character
	value, err := parser.Parse("121\n")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if value != 121 {
		t.Fatalf("expected 121, got %d", value)
	}
}

// TestInputParserDoesNotKeepState verifies parser is stateless
func TestInputParserDoesNotKeepState(t *testing.T) {
	parser := pipeline.NewInputParser()

	// Parse first value
	_, _ = parser.Parse("145")
	// Parse second value - should not depend on first
	_, err := parser.Parse("110")

	if err != nil {
		t.Fatalf("parser should not depend on previous input")
	}
}

// TestInputParserTableDriven runs multiple test cases in a table-driven approach
func TestInputParserTableDriven(t *testing.T) {
	parser := pipeline.NewInputParser()

	// Define test cases
	cases := []struct {
		name   string // Test case name
		input  string // Input string to parse
		expect int    // Expected parsed value
		hasErr bool   // Whether error is expected
	}{
		{"valid", "200", 200, false},
		{"spaces", " 99 ", 99, false},
		{"newline", "42\n", 42, false},
		{"letters", "foo", 0, true},
		{"mixed", "12a", 0, true},
		{"empty", "", 0, true},
	}

	// Run each test case
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Parse the input
			value, err := parser.Parse(tc.input)

			// Verify error expectation
			if tc.hasErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tc.hasErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			// Verify parsed value
			if !tc.hasErr && value != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, value)
			}
		})
	}
}
