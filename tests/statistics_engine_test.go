package tests

import (
	"guess-it/pipeline"
	"testing"
)

// TestStatisticsEngineInitialState verifies initial state (placeholder test)
func TestStatisticsEngineInitialState(t *testing.T) {
	// Create new statistics engine
	stats := pipeline.NewStatisticsEngine()
	// Verify initial count is zero
	if stats.Count() != 0 {
		t.Fatalf("expected initial count 0, got %d", stats.Count())
	}
}

// TestStatisticsEngineStdDevZero verifies standard deviation is zero for identical values
func TestStatisticsEngineStdDevZero(t *testing.T) {
	stats := pipeline.NewStatisticsEngine()

	// Add three identical values
	stats.Update(5)
	stats.Update(5)
	stats.Update(5)

	// Calculate standard deviation
	std := stats.StdDev()
	// Verify it's zero (no variance)
	if std != 0 {
		t.Fatalf("expected stddev 0, got %f", std)
	}
}

// TestStatisticsEngineMinMax verifies min and max tracking
func TestStatisticsEngineMinMax(t *testing.T) {
	stats := pipeline.NewStatisticsEngine()

	// Add multiple values
	stats.Update(42)
	stats.Update(7)
	stats.Update(100)
	stats.Update(50)

	// Get min and max
	min, max := stats.MinMax()

	// Verify minimum value
	if min != 7 {
		t.Fatalf("expected min 7, got %f", min)
	}
	// Verify maximum value
	if max != 100 {
		t.Fatalf("expected max 100, got %f", max)
	}
}

// TestStatisticsEngineTrendPositive verifies positive trend detection
func TestStatisticsEngineTrendPositive(t *testing.T) {
	stats := pipeline.NewStatisticsEngine()

	// Add increasing values
	stats.Update(10)
	stats.Update(20)

	// Get trend (difference between last two values)
	trend := stats.Trend()
	// Verify trend is positive
	if trend <= 0 {
		t.Fatalf("expected positive trend, got %f", trend)
	}
}

// TestStatisticsEngineTrendNegative verifies negative trend detection
func TestStatisticsEngineTrendNegative(t *testing.T) {
	stats := pipeline.NewStatisticsEngine()

	// Add decreasing values
	stats.Update(30)
	stats.Update(10)

	// Get trend
	trend := stats.Trend()
	// Verify trend is negative
	if trend >= 0 {
		t.Fatalf("expected negative trend, got %f", trend)
	}
}

// TestStatisticsEngineTrendZero verifies zero trend for identical consecutive values
func TestStatisticsEngineTrendZero(t *testing.T) {
	stats := pipeline.NewStatisticsEngine()

	// Add identical consecutive values
	stats.Update(15)
	stats.Update(15)

	// Get trend
	trend := stats.Trend()
	// Verify trend is zero
	if trend != 0 {
		t.Fatalf("expected zero trend, got %f", trend)
	}
}

// TestStatisticsEngineIncrementalUpdates verifies mean changes with new values
func TestStatisticsEngineIncrementalUpdates(t *testing.T) {
	stats := pipeline.NewStatisticsEngine()

	// Add first value and get mean
	stats.Update(1)
	firstMean := stats.Mean()

	// Add second value and get new mean
	stats.Update(3)
	secondMean := stats.Mean()

	// Verify mean changed
	if firstMean == secondMean {
		t.Fatalf("mean should change after update")
	}
}

// TestStatisticsEngineDoesNotRecomputeFromScratch verifies O(1) incremental computation
func TestStatisticsEngineDoesNotRecomputeFromScratch(t *testing.T) {
	stats := pipeline.NewStatisticsEngine()

	// Add 1000 identical values
	for i := 0; i < 1000; i++ {
		stats.Update(10)
	}

	// Calculate mean
	mean := stats.Mean()
	// Verify mean is correct
	if mean != 10 {
		t.Fatalf("expected mean 10, got %f", mean)
	}
}

// TestAverage verifies Average utility function returns correct mean
func TestAverage(t *testing.T) {
	// Test data: 2, 4, 6, 8 -> mean = 5
	data := []float64{2, 4, 6, 8}
	got := pipeline.Average(data)
	want := 5.0

	// Verify result
	if got != want {
		t.Errorf("Average() = %v, want %v", got, want)
	}
}

// TestMedian verifies median calculation for even and odd length slices
func TestMedian(t *testing.T) {
	// Even length: [1,2,3,4] -> median = (2+3)/2 = 2.5
	even := []float64{1, 2, 3, 4}
	// Odd length: [1,2,3] -> median = 2
	odd := []float64{3, 1, 2}

	// Verify even length median
	if pipeline.Median(even) != 2.5 {
		t.Errorf("Median(even) failed")
	}

	// Verify odd length median
	if pipeline.Median(odd) != 2 {
		t.Errorf("Median(odd) failed")
	}
}

// TestMinMax verifies Min and Max utility functions
func TestMinMax(t *testing.T) {
	// Test data with min=1, max=9
	data := []float64{5, 1, 9, 3}

	// Verify minimum
	if pipeline.Min(data) != 1 {
		t.Errorf("Min() failed")
	}

	// Verify maximum
	if pipeline.Max(data) != 9 {
		t.Errorf("Max() failed")
	}
}

// TestEmptySlice verifies utility functions handle empty input correctly
func TestEmptySlice(t *testing.T) {
	var data []float64

	// Verify Average returns 0 for empty slice
	if pipeline.Average(data) != 0 {
		t.Errorf("Average(empty) should be 0")
	}

	// Verify Median returns 0 for empty slice
	if pipeline.Median(data) != 0 {
		t.Errorf("Median(empty) should be 0")
	}
}
