
# CLI Application Guide

This CLI application is written in Go and is designed to analyze summary statistics in batches from a given CSV data file.

## Roles of the Program and Data

**Program:** The Go program reads data from a CSV file, processes it, and writes the summary statistics to an output file. The statistics performed in this app are the minimum, standard deviation, mean, median, and max. The program uses command-line flags to specify the input and output files.

**Data:** The data must be in a CSV file with column names and numerical values. The program reads this data, calculates summary statistics, and writes the results to the output file defined be the user.

# Use Instructions
1. Open a bash terminal

2. Navigate to the directory that holds the statscli executable from the command line. 

3. The bash command to run the script is as follows:

./statscli -input "path\to\dataset\theDataset.csv" -output "path\to\dataset\dataOutput.txt"

It is recomended that the user has the app executable, the data file and the output in the same directory to avoid adding a path, but is not nessecary.

For example: ./statscli -input "housesInput.csv" -output "dataOutput.txt" would all be in the same directory and run without issue.

# Summary for Managment

Go has proven viable for  the development and deployment of the CLI application for analyzing summary statistics. It is recomended that managment proceed with the switch to Go for these batch tasks. Go demonstrated to be the fastest of the three programming languaugs as seen in execution_times.txt executing at 2.67seconds when CPU times was run with time ./statsCLI command at the terminal.