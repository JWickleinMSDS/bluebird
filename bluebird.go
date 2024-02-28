// bluebird.go

package bluebird

import (
	"sort"
)

// TrimmedMean calculates the trimmed mean of a slice of integers or floats
func TrimmedMean(data []float64, lowerTrim, upperTrim float64) float64 {
	// Sort data now
	sortedData := make([]float64, len(data))
	copy(sortedData, data)
	sort.Float64s(sortedData)

	// Calculate the number of things to be trimmed
	lowerTrimCount := int(float64(len(sortedData)) * lowerTrim)
	upperTrimCount := int(float64(len(sortedData)) * upperTrim)

	// Trim the data now
	trimmedData := sortedData[lowerTrimCount : len(sortedData)-upperTrimCount]

	// Calculate the mean of trimmed data
	sum := 0.0
	for _, val := range trimmedData {
		sum += val
	}
	return sum / float64(len(trimmedData))
}
