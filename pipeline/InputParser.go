package pipeline

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

var ErrInvalidInput = errors.New("invalid input")

type InputParser struct{}

func NewInputParser() *InputParser {
	return &InputParser{}
}

// Parse μετατρέπει raw string σε int, αγνοεί whitespace/newline
func (p *InputParser) Parse(raw string) (int, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return 0, ErrInvalidInput
	}
	n, err := strconv.Atoi(raw)
	if err != nil {
		return 0, ErrInvalidInput
	}
	return n, nil
}

// readInput διαβάζει μια γραμμή από stdin και την parse-άρει
func readInput() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	parser := NewInputParser()
	return parser.Parse(line)
}
