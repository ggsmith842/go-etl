package targets

import (
	"encoding/csv"
	"errors"
	"os"
	"testing"
)

func TestCsvTgt(t *testing.T) {
	// Create test file
	file, err := os.Create("testtgt.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	w.WriteAll([][]string{
		{"Field1", "Field2", "Field3"},
		{"A", "B", "C"},
		{"D", "E", "F"},
	})
	w.Flush()

	// Open the test CSV file
	if _, err := os.Stat("testtgt.csv"); errors.Is(err, os.ErrNotExist) {
		t.Errorf("file not written")
	} else {
		err = os.Remove("testtgt.csv")
		if err != nil {
			t.Fatal(err)
		}

	}

}
