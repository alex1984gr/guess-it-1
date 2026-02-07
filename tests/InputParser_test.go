package tests

import (
	"errors"
	"guess-it/pipeline"
	"strings"
	"testing"
)

// The Input Parser is responsible ONLY for parsing raw input strings
// into integers. It must not keep state.

func TestInputParserValidInteger(t *testing.T) {
	parser := pipeline.NewInputParser()

	value, err := parser.Parse("189")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if value != 189 {
		t.Fatalf("expected 189, got %d", value)
	}
}

func TestInputParserWithWhitespace(t *testing.T) {
	parser := pipeline.NewInputParser()

	value, err := parser.Parse("  113  ")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if value != 113 {
		t.Fatalf("expected 113, got %d", value)
	}
}

func TestInputParserInvalidInput(t *testing.T) {
	parser := pipeline.NewInputParser()

	_, err := parser.Parse("abc")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestInputParserEmptyInput(t *testing.T) {
	parser := pipeline.NewInputParser()

	_, err := parser.Parse("")

	if err == nil {
		t.Fatalf("expected error for empty input")
	}
}

func TestInputParserNewlineInput(t *testing.T) {
	parser := pipeline.NewInputParser()

	value, err := parser.Parse("121\n")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if value != 121 {
		t.Fatalf("expected 121, got %d", value)
	}
}

func TestInputParserDoesNotKeepState(t *testing.T) {
	parser := pipeline.NewInputParser()

	_, _ = parser.Parse("145")
	_, err := parser.Parse("110")

	if err != nil {
		t.Fatalf("parser should not depend on previous input")
	}
}

func TestInputParserReturnsTypedError(t *testing.T) {
	parser := pipeline.NewInputParser()

	_, err := parser.Parse("xyz")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !errors.Is(err, pipeline.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

// Table-driven test for robustness
func TestInputParserTableDriven(t *testing.T) {
	parser := pipeline.NewInputParser()

	cases := []struct {
		name   string
		input  string
		expect int
		hasErr bool
	}{
		{"valid", "200", 200, false},
		{"spaces", " 99 ", 99, false},
		{"newline", "42\n", 42, false},
		{"letters", "foo", 0, true},
		{"mixed", "12a", 0, true},
		{"empty", "", 0, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			value, err := parser.Parse(strings.TrimSpace(tc.input))

			if tc.hasErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tc.hasErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.hasErr && value != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, value)
			}
		})
	}
}
