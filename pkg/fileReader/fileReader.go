package fileReader

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Read(path string) ([][]string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return records, nil
}
