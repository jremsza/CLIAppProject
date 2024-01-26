package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/montanaflynn/stats"
)

func processFile(filename string) {

	//open CSV file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}

	defer file.Close()

	//create a new reader
	r := csv.NewReader(file)

	//read in headers the CSV records
	headers, err := r.Read()
	if err != nil {
		fmt.Println("Error reading in headers:", err)
		return
	}

	//slice of slices to hold data
	var data [][]float64 = make([][]float64, len(headers))

	// read in the rest of the records
	for {
		record, err := r.Read()
		if err != nil {
			break
		}

		//parse each of the values in the record into a float64
		for i, value := range record {
			num, err := strconv.ParseFloat(value, 64)
			if err != nil {
				fmt.Printf("Error parsing value in column %d: %v\n:", i, err)
				continue
			}

			//append the float64 value to the appropriate slice
			data[i] = append(data[i], num)
		}
	}
	//create a new file for output - include flag to append to the file if it already exists
	output, err := os.OpenFile("housesOutputGo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening output file:", err)
		return
	}

	defer output.Close()

	// Compute and print statistics for each column
	for i, columnData := range data {
		min, _ := stats.Min(columnData)
		sd, _ := stats.StandardDeviation(columnData)
		mean, _ := stats.Mean(columnData)
		median, _ := stats.Median(columnData)
		max, _ := stats.Max(columnData)

		//output the results
		outputString := fmt.Sprintf("%s\nMinimum: %.2f\nStDev: %.2f\nMean: %.2f\nMedian: %.2f\nMaximum: %.2f\n\n", headers[i], min, sd, mean, median, max)
		_, err := output.WriteString(outputString)
		if err != nil {
			fmt.Println("Error writing output:", err)
			return
		}
	}
}

func main() {

	filename := "housesInput.csv"

	// Run processFile 100 times
	for i := 0; i < 100; i++ {
		processFile(filename)
	}
}
