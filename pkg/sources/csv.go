package sources

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Record struct {
	DATE         string
	MORTGAGE30US float64
}

func CsvSrc(dir string) {
	file, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	//read csv values
	csvReader := csv.NewReader(file)

	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i, row := range data {
		if i == 0 {
			continue
		}

		rate, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		record := Record{
			DATE:         row[0],
			MORTGAGE30US: rate,
		}
		fmt.Println(record.DATE, record.MORTGAGE30US)
	}

}
