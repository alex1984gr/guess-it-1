package pipeline

import (
	"math"
)

// StatisticsEngine υπολογίζει mean, stddev, min, max, trend
type StatisticsEngine struct {
	count int
	sum   float64
	sumSq float64
	minV  float64
	maxV  float64
	last  float64
	trend float64
	init  bool
}

func NewStatisticsEngine() *StatisticsEngine {
	return &StatisticsEngine{
		minV: math.Inf(1),
		maxV: math.Inf(-1),
	}
}

// Update προσθέτει νέο αριθμό
func (s *StatisticsEngine) Update(value int) {
	x := float64(value)

	if s.init {
		s.trend = x - s.last
	} else {
		s.init = true
		s.trend = 0
	}
	s.last = x

	s.count++
	s.sum += x
	s.sumSq += x * x

	if x < s.minV {
		s.minV = x
	}
	if x > s.maxV {
		s.maxV = x
	}
}

// Count επιστρέφει αριθμό στοιχείων
func (s *StatisticsEngine) Count() int {
	return s.count
}

// Mean υπολογίζει μέσο όρο
func (s *StatisticsEngine) Mean() float64 {
	if s.count == 0 {
		return 0
	}
	return s.sum / float64(s.count)
}

// StdDev υπολογίζει standard deviation
func (s *StatisticsEngine) StdDev() float64 {
	if s.count == 0 {
		return 0
	}
	mean := s.Mean()
	variance := (s.sumSq / float64(s.count)) - mean*mean
	return math.Sqrt(variance)
}

// MinMax επιστρέφει min και max
func (s *StatisticsEngine) MinMax() (float64, float64) {
	if s.count == 0 {
		return 0, 0
	}
	return s.minV, s.maxV
}

// Trend επιστρέφει τελευταίο trend
func (s *StatisticsEngine) Trend() float64 {
	return s.trend
}
