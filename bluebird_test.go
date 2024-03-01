package bluebird_test

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/JWickleinMSDS/bluebird"
)

func TestTrimmedMean(t *testing.T) {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	// Generate random samples of integers
	integerData := generateRandomIntegers(101, 1, 100)
	// Calculate the trimmed mean
	goTrimmedMean := bluebird.TrimmedMean(integerData, 0.05, 0.05)
	// Calculate the trimmed mean in R
	rTrimmedMean, err := calculateRTrimmedMean(integerData)
	if err != nil {
		t.Errorf("error calculating R trimmed mean: %v", err)
		return
	}

	// Round the trimmed means to three decimal places so the comparison has some slack
	goTrimmedMeanRound := fmt.Sprintf("%.3f", goTrimmedMean)
	rTrimmedMeanRound := fmt.Sprintf("%.3f", rTrimmedMean)

	// Compare the results
	if goTrimmedMeanRound != rTrimmedMeanRound {
		t.Errorf("trimmed means do not match: Go=%s, R=%s", goTrimmedMeanRound, rTrimmedMeanRound)
	}

	// Generate random samples of floats
	floatData := generateRandomFloats(101, 1.0, 100.0)
	// Calculate the trimmed mean
	goTrimmedMeanFloat := bluebird.TrimmedMean(floatData, 0.05, 0.05)
	// Calculate the trimmed mean using R
	rTrimmedMeanFloat, err := calculateRTrimmedMean(floatData)
	if err != nil {
		t.Errorf("error calculating R trimmed mean for float data: %v", err)
		return
	}

	// Round the trimmed means so you have some slack
	goTrimmedMeanFloatRound := fmt.Sprintf("%.3f", goTrimmedMeanFloat)
	rTrimmedMeanFloatRound := fmt.Sprintf("%.3f", rTrimmedMeanFloat)

	// Compare the results
	if goTrimmedMeanFloatRound != rTrimmedMeanFloatRound {
		t.Errorf("trimmed means for float data do not match: Go=%s, R=%s", goTrimmedMeanFloatRound, rTrimmedMeanFloatRound)
	}
}

// Generate random floats
func generateRandomFloats(n int, min, max float64) []float64 {
	data := make([]float64, n)
	for i := 0; i < n; i++ {
		data[i] = min + rand.Float64()*(max-min)
	}
	return data
}

// Generate random samples of integers
func generateRandomIntegers(n int, min, max int) []float64 {
	data := make([]float64, n)
	for i := 0; i < n; i++ {
		data[i] = float64(rand.Intn(max-min+1) + min)
	}
	return data
}

// Calculate trimmed mean using R script
func calculateRTrimmedMean(data []float64) (float64, error) {
	cmd := exec.Command("Rscript", "bluebird.R")

	// Convert data to string slice
	dataStrings := make([]string, len(data))
	for i, val := range data {
		dataStrings[i] = strconv.FormatFloat(val, 'f', -1, 64)
	}

	// Append data as command-line arguments
	cmd.Args = append(cmd.Args, dataStrings...)

	// Capture output
	out, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("error executing R script: %v", err)
	}

	// Parse R
	rTrimmedMean, err := strconv.ParseFloat(strings.TrimSpace(string(out)), 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing R output: %v", err)
	}

	return rTrimmedMean, nil
}
