package pipeline

// PredictionStrategy generates prediction ranges based on statistics
type PredictionStrategy struct {
	multiplier float64 // Controls the width of the prediction range
}

// NewPredictionStrategy creates a new PredictionStrategy with given multiplier
func NewPredictionStrategy(mult float64) *PredictionStrategy {
	return &PredictionStrategy{multiplier: mult}
}

// Predict returns a range around mean + trend
func (ps *PredictionStrategy) Predict(mean, stddev, trend float64) (float64, float64) {
	// Center the prediction around mean adjusted by trend
	center := mean + trend
	// Calculate range width based on standard deviation
	width := stddev * ps.multiplier
	if width == 0 {
		// Minimum range to avoid zero-width predictions
		width = 1
	}
	// Return lower and upper bounds
	return center - width, center + width
}
