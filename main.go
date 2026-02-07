package main

import (
	"bufio"   // For reading input from stdin
	"fmt"     // For printing output
	"os"      // For accessing stdin
	"strconv" // For converting string to int
	"strings" // For trimming whitespace

	"guess-it/pipeline" // Our custom pipeline package
)

func main() {
	// Create a new StatisticsEngine to track numbers
	stats := pipeline.NewStatisticsEngine()

	// Create a buffered reader to read lines from stdin
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter numbers one by one. Press Ctrl+C to exit.")

	// Infinite loop to continuously read numbers
	for {
		fmt.Print("> ") // Prompt for user input

		// Read a line from stdin
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Remove leading/trailing whitespace
		line = strings.TrimSpace(line)
		if line == "" {
			continue // Skip empty lines
		}

		// Convert the input string to an integer
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Invalid number:", line)
			continue
		}

		// Add the number to the statistics engine
		stats.Update(num)

		// Calculate mean and standard deviation
		mean := stats.Mean()
		std := stats.StdDev()

		// Predict the next number's range: mean ± stddev
		lower := mean - std
		upper := mean + std

		// Round to integer for display
		fmt.Printf("Predicted range for next number: %d %d\n", int(lower), int(upper))
	}
}
