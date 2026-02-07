package tests

import (
	"guess-it/pipeline"
	"testing"
)

// The Prediction Strategy takes mean, stddev, trend and outputs a predicted range

func TestPredictionRangeValidity(t *testing.T) {
	strategy := pipeline.NewPredictionStrategy(1.0)

	lower, upper := strategy.Predict(100, 10, 5)

	if lower >= upper {
		t.Fatalf("invalid range: %f %f", lower, upper)
	}
}

func TestPredictionCenteredAroundMeanPlusTrend(t *testing.T) {
	strategy := pipeline.NewPredictionStrategy(1.0)

	mean := 50.0
	std := 10.0
	trend := 5.0

	lower, upper := strategy.Predict(mean, std, trend)
	center := (lower + upper) / 2

	expected := mean + trend
	if center != expected {
		t.Fatalf("expected center %f, got %f", expected, center)
	}
}

func TestPredictionNarrowerWithLowerStd(t *testing.T) {
	strategy := pipeline.NewPredictionStrategy(1.0)

	l1, u1 := strategy.Predict(100, 20, 0)
	l2, u2 := strategy.Predict(100, 5, 0)

	if (u2 - l2) >= (u1 - l1) {
		t.Fatalf("expected narrower range with lower stddev")
	}
}

func TestPredictionAdjustsForTrend(t *testing.T) {
	strategy := pipeline.NewPredictionStrategy(1.0)

	mean := 100.0
	std := 10.0
	trend := -20.0

	lower, upper := strategy.Predict(mean, std, trend)

	center := (lower + upper) / 2
	if center != mean+trend {
		t.Fatalf("expected center %f, got %f", mean+trend, center)
	}
}

func TestPredictionHandlesZeroStd(t *testing.T) {
	strategy := pipeline.NewPredictionStrategy(1.0)

	lower, upper := strategy.Predict(50, 0, 5)
	if lower >= upper {
		t.Fatalf("expected valid range even with zero std: %f %f", lower, upper)
	}
}

func TestPredictionConsistency(t *testing.T) {
	strategy := pipeline.NewPredictionStrategy(1.0)

	mean := 100.0
	std := 15.0
	trend := 0.0

	l1, u1 := strategy.Predict(mean, std, trend)
	l2, u2 := strategy.Predict(mean, std, trend)

	if l1 != l2 || u1 != u2 {
		t.Fatalf("expected deterministic output for same inputs")
	}
}
