package pipeline

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// InputParser handles parsing of raw string input into integers
type InputParser struct{}

// NewInputParser creates a new InputParser instance
func NewInputParser() *InputParser {
	return &InputParser{}
}

// Parse converts a raw string to an integer after trimming whitespace
func (p *InputParser) Parse(raw string) (int, error) {
	// Remove leading and trailing whitespace
	raw = strings.TrimSpace(raw)
	// Convert string to integer
	return strconv.Atoi(raw)
}

// ReadInputFrom reads from any io.Reader (for mock stdin in tests)
func ReadInputFrom(r io.Reader) (int, error) {
	// Create a buffered reader for efficient reading
	reader := bufio.NewReader(r)
	// Read until newline character
	line, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	// Create parser and parse the line
	parser := NewInputParser()
	return parser.Parse(line)
}
