package pipeline

import (
	"math"
)

// StatisticsEngine computes mean, stddev, min, max, and trend incrementally
type StatisticsEngine struct {
	count int     // Number of values processed
	sum   float64 // Sum of all values
	sumSq float64 // Sum of squared values (for variance calculation)
	minV  float64 // Minimum value seen
	maxV  float64 // Maximum value seen
	last  float64 // Last value received
	trend float64 // Difference between current and previous value
	init  bool    // Whether we have received at least one value
}

// NewStatisticsEngine creates a new StatisticsEngine instance
func NewStatisticsEngine() *StatisticsEngine {
	return &StatisticsEngine{
		minV: math.Inf(1),  // Initialize to positive infinity
		maxV: math.Inf(-1), // Initialize to negative infinity
	}
}

// Update adds a new value and updates all statistics
func (s *StatisticsEngine) Update(value int) {
	x := float64(value)

	// Calculate trend as difference from last value
	if s.init {
		s.trend = x - s.last
	} else {
		// First value has no trend
		s.init = true
		s.trend = 0
	}
	s.last = x

	// Update count and sums
	s.count++
	s.sum += x
	s.sumSq += x * x

	// Update min and max
	if x < s.minV {
		s.minV = x
	}
	if x > s.maxV {
		s.maxV = x
	}
}

// Count returns the number of values processed
func (s *StatisticsEngine) Count() int {
	return s.count
}

// Mean calculates the average of all values
func (s *StatisticsEngine) Mean() float64 {
	if s.count == 0 {
		return 0
	}
	return s.sum / float64(s.count)
}

// StdDev calculates the standard deviation
func (s *StatisticsEngine) StdDev() float64 {
	if s.count == 0 {
		return 0
	}
	mean := s.Mean()
	// Variance = E[X²] - E[X]²
	variance := (s.sumSq / float64(s.count)) - mean*mean
	return math.Sqrt(variance)
}

// MinMax returns the minimum and maximum values
func (s *StatisticsEngine) MinMax() (float64, float64) {
	if s.count == 0 {
		return 0, 0
	}
	return s.minV, s.maxV
}

// Trend returns the most recent trend (difference from previous value)
func (s *StatisticsEngine) Trend() float64 {
	return s.trend
}
