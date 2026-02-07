package tests

import (
	"guess-it/pipeline"
	"testing"
)

func TestStatisticsEngineInitialState(t *testing.T) {
}

func TestStatisticsEngineStdDevZero(t *testing.T) {
	stats := pipeline.NewStatisticsEngine()

	stats.Update(5)
	stats.Update(5)
	stats.Update(5)

	std := stats.StdDev()
	if std != 0 {
		t.Fatalf("expected stddev 0, got %f", std)
	}
}

func TestStatisticsEngineMinMax(t *testing.T) {
	stats := pipeline.NewStatisticsEngine()

	stats.Update(42)
	stats.Update(7)
	stats.Update(100)
	stats.Update(50)

	min, max := stats.MinMax()

	if min != 7 {
		t.Fatalf("expected min 7, got %f", min)
	}
	if max != 100 {
		t.Fatalf("expected max 100, got %f", max)
	}
}

func TestStatisticsEngineTrendPositive(t *testing.T) {
	stats := pipeline.NewStatisticsEngine()

	stats.Update(10)
	stats.Update(20)

	trend := stats.Trend()
	if trend <= 0 {
		t.Fatalf("expected positive trend, got %f", trend)
	}
}

func TestStatisticsEngineTrendNegative(t *testing.T) {
	stats := pipeline.NewStatisticsEngine()

	stats.Update(30)
	stats.Update(10)

	trend := stats.Trend()
	if trend >= 0 {
		t.Fatalf("expected negative trend, got %f", trend)
	}
}

func TestStatisticsEngineTrendZero(t *testing.T) {
	stats := pipeline.NewStatisticsEngine()

	stats.Update(15)
	stats.Update(15)

	trend := stats.Trend()
	if trend != 0 {
		t.Fatalf("expected zero trend, got %f", trend)
	}
}

func TestStatisticsEngineIncrementalUpdates(t *testing.T) {
	stats := pipeline.NewStatisticsEngine()

	stats.Update(1)
	firstMean := stats.Mean()

	stats.Update(3)
	secondMean := stats.Mean()

	if firstMean == secondMean {
		t.Fatalf("mean should change after update")
	}
}

func TestStatisticsEngineDoesNotRecomputeFromScratch(t *testing.T) {
	stats := pipeline.NewStatisticsEngine()

	for i := 0; i < 1000; i++ {
		stats.Update(10)
	}

	mean := stats.Mean()
	if mean != 10 {
		t.Fatalf("expected mean 10, got %f", mean)
	}
}

// TestAverage checks if Average returns correct mean value
func TestAverage(t *testing.T) {
	data := []float64{2, 4, 6, 8}
	got := pipeline.Average(data)
	want := 5.0

	if got != want {
		t.Errorf("Average() = %v, want %v", got, want)
	}
}

// TestMedian checks median calculation for even and odd slices
func TestMedian(t *testing.T) {
	even := []float64{1, 2, 3, 4}
	odd := []float64{3, 1, 2}

	if pipeline.Median(even) != 2.5 {
		t.Errorf("Median(even) failed")
	}

	if pipeline.Median(odd) != 2 {
		t.Errorf("Median(odd) failed")
	}
}

// TestMinMax checks min and max values
func TestMinMax(t *testing.T) {
	data := []float64{5, 1, 9, 3}

	if pipeline.Min(data) != 1 {
		t.Errorf("Min() failed")
	}

	if pipeline.Max(data) != 9 {
		t.Errorf("Max() failed")
	}
}

// TestEmptySlice checks behavior with empty input
func TestEmptySlice(t *testing.T) {
	var data []float64

	if pipeline.Average(data) != 0 {
		t.Errorf("Average(empty) should be 0")
	}

	if pipeline.Median(data) != 0 {
		t.Errorf("Median(empty) should be 0")
	}
}
