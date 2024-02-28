// bluebird_test.go

package bluebird_test

import (
	"math"
	"testing"

	"github.com/JWickleinMSDS/bluebird"
	"github.com/stretchr/testify/assert"
)

func TestTrimmedMean(t *testing.T) {
	// Test with integers
	integerData := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := bluebird.TrimmedMean(integerData, 0.05, 0.05)
	assert.Equal(t, 5.5, result)

	// Test with floats
	floatData := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.0}
	result = bluebird.TrimmedMean(floatData, 0.05, 0.05)
	assert.Equal(t, 5.775, math.Round(result*1000)/1000) // Rounding to facilitate comparison
}
