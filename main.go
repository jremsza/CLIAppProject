package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/montanaflynn/stats"
)

func main() {
	//start timing
	start := time.Now()

	//open CSV file
	file, err := os.Open("housesInput.csv")
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
	//create a new file for output
	output, err := os.Create("housesOutputGo.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer output.Close()

	// Compute and print statistics for each column
	for i, columnData := range data {
		min, _ := stats.Min(columnData)
		max, _ := stats.Max(columnData)
		mean, _ := stats.Mean(columnData)
		median, _ := stats.Median(columnData)

		//fmt.Printf("Column %d: min=%v, max=%v, mean=%v, median=%v\n", i+1, min, max, mean, median)

		//output the results
		outputString := fmt.Sprintf("%s\nMinimum: %.2f\nMaximum: %.2f\nMean: %.2f\nMedian: %.2f\n\n", headers[i], min, max, mean, median)
		_, err := output.WriteString(outputString)
		if err != nil {
			fmt.Println("Error writing output:", err)
			return
		}
	}
	//calculate and print the elapsed time
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}
