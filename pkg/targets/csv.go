package targets

import (
	"encoding/csv"
	"fmt"
	d "goetl/data"
	"os"
	"reflect"
)

func CsvTgt(name string, rows []d.Record) {

	// Create new csv file
	file, err := os.Create(name + ".csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write Headers
	var r d.Record
	t := reflect.TypeOf(r)
	headers := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		headers[i] = t.Field(i).Name
	}
	writer.Write(headers)

	// Create a slice to hold the records
	data := make([][]string, 0, len(rows))
	for _, row := range rows {
		r := []string{row.DATE, fmt.Sprint(row.MORTGAGE30US)}
		data = append(data, r)
	}

	// Write all records to the CSV file
	err = writer.WriteAll(data)
	if err != nil {
		panic(err)
	}

}
