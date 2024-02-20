package main

import (
	"bufio"
	"encoding/csv"
	"os"
)

func readCSVFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func readTextFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// func main() {
// 	// Example for reading a CSV file
// 	csvFilename := "example.csv"
// 	csvData, err := readCSVFile(csvFilename)
// 	if err != nil {
// 		fmt.Println("Error reading CSV file:", err)
// 		return
// 	}
// 	fmt.Println("CSV Data:")
// 	fmt.Println(csvData)

// 	// Example for reading a text (txt) file
// 	txtFilename := "example.txt"
// 	txtData, err := readTextFile(txtFilename)
// 	if err != nil {
// 		fmt.Println("Error reading text file:", err)
// 		return
// 	}
// 	fmt.Println("Text Data:")
// 	fmt.Println(txtData)
// }
