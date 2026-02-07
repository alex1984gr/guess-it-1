package tests

import (
	"guess-it/pipeline"
	"testing"
)

// TestPredictionRangeValidity verifies lower bound is less than upper bound
func TestPredictionRangeValidity(t *testing.T) {
	// Create strategy with multiplier 1.0
	strategy := pipeline.NewPredictionStrategy(1.0)

	// Generate prediction: mean=100, stddev=10, trend=5
	lower, upper := strategy.Predict(100, 10, 5)

	// Verify valid range (lower < upper)
	if lower >= upper {
		t.Fatalf("invalid range: %f %f", lower, upper)
	}
}

// TestPredictionCenteredAroundMeanPlusTrend verifies prediction center calculation
func TestPredictionCenteredAroundMeanPlusTrend(t *testing.T) {
	strategy := pipeline.NewPredictionStrategy(1.0)

	// Set test parameters
	mean := 50.0
	std := 10.0
	trend := 5.0

	// Generate prediction
	lower, upper := strategy.Predict(mean, std, trend)
	// Calculate center of range
	center := (lower + upper) / 2

	// Verify center equals mean + trend
	expected := mean + trend
	if center != expected {
		t.Fatalf("expected center %f, got %f", expected, center)
	}
}

// TestPredictionNarrowerWithLowerStd verifies range width decreases with lower stddev
func TestPredictionNarrowerWithLowerStd(t *testing.T) {
	strategy := pipeline.NewPredictionStrategy(1.0)

	// Generate prediction with high stddev
	l1, u1 := strategy.Predict(100, 20, 0)
	// Generate prediction with low stddev
	l2, u2 := strategy.Predict(100, 5, 0)

	// Verify second range is narrower
	if (u2 - l2) >= (u1 - l1) {
		t.Fatalf("expected narrower range with lower stddev")
	}
}

// TestPredictionAdjustsForTrend verifies trend shifts the prediction center
func TestPredictionAdjustsForTrend(t *testing.T) {
	strategy := pipeline.NewPredictionStrategy(1.0)

	// Set parameters with negative trend
	mean := 100.0
	std := 10.0
	trend := -20.0

	// Generate prediction
	lower, upper := strategy.Predict(mean, std, trend)

	// Calculate center
	center := (lower + upper) / 2
	// Verify center adjusted by trend
	if center != mean+trend {
		t.Fatalf("expected center %f, got %f", mean+trend, center)
	}
}

// TestPredictionHandlesZeroStd verifies minimum range when stddev is zero
func TestPredictionHandlesZeroStd(t *testing.T) {
	strategy := pipeline.NewPredictionStrategy(1.0)

	// Generate prediction with zero stddev
	lower, upper := strategy.Predict(50, 0, 5)
	// Verify valid range (minimum width of 1 is applied)
	if lower >= upper {
		t.Fatalf("expected valid range even with zero std: %f %f", lower, upper)
	}
}

// TestPredictionConsistency verifies deterministic output for same inputs
func TestPredictionConsistency(t *testing.T) {
	strategy := pipeline.NewPredictionStrategy(1.0)

	// Set test parameters
	mean := 100.0
	std := 15.0
	trend := 0.0

	// Generate two predictions with same inputs
	l1, u1 := strategy.Predict(mean, std, trend)
	l2, u2 := strategy.Predict(mean, std, trend)

	// Verify outputs are identical
	if l1 != l2 || u1 != u2 {
		t.Fatalf("expected deterministic output for same inputs")
	}
}
