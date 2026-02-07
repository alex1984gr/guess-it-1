package pipeline

import "sort"

// Average returns the mean of a slice of float64 values
func Average(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	sum := 0.0
	// Sum all values
	for _, v := range data {
		sum += v
	}
	// Return mean
	return sum / float64(len(data))
}

// Median returns the median value of a slice
func Median(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	// Create a copy to avoid modifying original
	sorted := make([]float64, len(data))
	copy(sorted, data)
	// Sort the copy
	sort.Float64s(sorted)
	n := len(sorted)
	// If even number of elements, return average of middle two
	if n%2 == 0 {
		return (sorted[n/2-1] + sorted[n/2]) / 2
	}
	// If odd, return middle element
	return sorted[n/2]
}

// Min returns the smallest value in a slice
func Min(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	min := data[0]
	// Find minimum value
	for _, v := range data {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the largest value in a slice
func Max(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	// Find maximum value
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}
