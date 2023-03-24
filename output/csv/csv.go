package csv

import (
	"encoding/csv"
	"os"
)

// Generate generates the CSV in the given path
func Generate(inputHeader []string, inputRows [][]string, outPath string) (bool, error) {
	csvFile, err := os.Create(outPath)

	csvInput := make([][]string, len(inputRows)+1)
	header := make([][]string, 0)
	header = append(header, inputHeader)

	csvInput = append(header, inputRows...)

	if err != nil {
		return false, err
	}

	csvWriter := csv.NewWriter(csvFile)
	writeErr := csvWriter.WriteAll(csvInput)
	if writeErr != nil {
		return false, err
	}

	csvWriter.Flush()
	csvFile.Close()
	return true, nil
}
