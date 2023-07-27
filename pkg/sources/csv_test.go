package sources_test

import (
	"encoding/csv"
	"os"
	"reflect"
	"testing"
)

func TestCsvSrc(t *testing.T) {
	// Create test file
	file, err := os.Create("test.csv")
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
	file, err = os.Open("test.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Read the data from the test CSV file
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	// Compare read data to expected data
	expected := [][]string{
		{"Field1", "Field2", "Field3"},
		{"A", "B", "C"},
		{"D", "E", "F"},
	}
	if !reflect.DeepEqual(records, expected) {
		t.Errorf("got %v, want %v", records, expected)
	}

	err = os.Remove("test.csv")
	if err != nil {
		t.Fatal(err)
	}
}
