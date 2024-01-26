package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestProcessFile(t *testing.T) {
	// Create a temporary test file
	file, err := os.CreateTemp("", "testfile.csv")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	// Write test data to the file
	data := []string{"1", "2", "3", "4", "5"}
	writer := csv.NewWriter(file)
	writer.Write(data)
	writer.Flush()

	// Close the file
	file.Close()

	// Call the processFile function
	processFile(file.Name())

	// Example assertion:
	expectedResult := 15.0
	actualResult := calculateSum(data)
	if actualResult != expectedResult {
		t.Errorf("Expected sum to be %f, but got %f", expectedResult, actualResult)
	}
}

func calculateSum(data []string) float64 {
	numbers := make([]float64, len(data))
	for i, str := range data {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0.0
		}
		numbers[i] = num
	}
	sum, _ := stats.Sum(numbers)
	return sum
}
