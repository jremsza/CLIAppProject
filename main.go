package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/montanaflynn/stats"
)

func processFile(filename string, output *os.File) {
	// Open CSV file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	// Create a new reader
	r := csv.NewReader(file)

	// Read in headers from the CSV records
	headers, err := r.Read()
	if err != nil {
		fmt.Println("Error reading headers:", err)
		return
	}

	// Slice of slices to hold data
	var data [][]float64 = make([][]float64, len(headers))

	// Read in the rest of the records
	for {
		record, err := r.Read()
		if err == csv.ErrFieldCount {
			// Handle expected error if the number of fields is different
			continue
		}
		if err != nil {
			// Check if error is EOF
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading record:", err)
			return
		}

		// Parse each value in the record into a float64
		for i, value := range record {
			num, err := strconv.ParseFloat(value, 64)
			if err != nil {
				fmt.Printf("Error parsing value in column %d: %v\n", i, err)
				continue
			}
			data[i] = append(data[i], num)
		}
	}

	// Compute and print statistics for each column
	for i, columnData := range data {
		min, _ := stats.Min(columnData)
		sd, _ := stats.StandardDeviation(columnData)
		mean, _ := stats.Mean(columnData)
		median, _ := stats.Median(columnData)
		max, _ := stats.Max(columnData)

		outputString := fmt.Sprintf("%s\nMinimum: %.2f\nStDev: %.2f\nMean: %.2f\nMedian: %.2f\nMaximum: %.2f\n\n", headers[i], min, sd, mean, median, max)
		_, err := output.WriteString(outputString)
		if err != nil {
			fmt.Println("Error writing output:", err)
			return
		}
	}
}

func main() {
	inputFilename := flag.String("input", "housesInput.csv", "The name of the input CSV file")
	outputFilename := flag.String("output", "housesOutputGo.txt", "The name of the output file")
	flag.Parse()

	// Check if the input file exists and is readable
	if _, err := os.Stat(*inputFilename); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Error: Input file does not exist.")
		} else {
			fmt.Println("Error: Unable to read input file. Unsure the file is .csv format.")
		}
		return
	}

	// Check if the output file is writable
	output, err := os.OpenFile(*outputFilename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error: Unable to write to output file.")
		return
	}

	defer output.Close()

	for i := 0; i < 100; i++ {
		processFile(*inputFilename, output)
	}
}
