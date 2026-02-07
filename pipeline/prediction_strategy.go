package pipeline

type PredictionStrategy struct {
	multiplier float64
}

func NewPredictionStrategy(mult float64) *PredictionStrategy {
	return &PredictionStrategy{multiplier: mult}
}

// Predict επιστρέφει ένα εύρος γύρω από mean + trend
func (ps *PredictionStrategy) Predict(mean, stddev, trend float64) (float64, float64) {
	center := mean + trend
	width := stddev * ps.multiplier
	if width == 0 {
		width = 1 // ελάχιστο range για να μην είναι 0
	}
	return center - width, center + width
}
