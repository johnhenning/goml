package utilities

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

// LoadCSVAsDataset takes a filename of a csv file and
func LoadCSVAsDataset(filepath string) ([][]string, error) {
	csvFile, _ := os.Open(filepath)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var dataset [][]string
	for {
		row, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		dataset = append(dataset, row)
	}
	return dataset, nil
}
