package bluebird

import (
	"sort"
)

// TrimmedMean calculates the trimmed mean of a slice of integers or floats
func TrimmedMean(data []float64, lowerTrim, upperTrim float64) float64 {
	// Sort the data!
	sortedData := make([]float64, len(data))
	copy(sortedData, data)
	sort.Float64s(sortedData)

	// Calculate the number of items to be trimmed
	lowerTrimCount := int(float64(len(sortedData)) * lowerTrim)
	upperTrimCount := int(float64(len(sortedData)) * upperTrim)

	// Trim the data!
	trimmedData := sortedData[lowerTrimCount : len(sortedData)-upperTrimCount]

	// Calculate the mean
	sum := 0.0
	for _, val := range trimmedData {
		sum += val
	}
	return sum / float64(len(trimmedData))
}
