package sources

import (
	"encoding/csv"
	// "goetl/data"
	// d "goetl/data"
	"log"
	"os"
)

func CsvSrc(dir string) [][]string {
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

	// rows := make([]d.Record, 0, len(data))

	// //loop through rows and store in map
	// for i, row := range data {
	// 	if i == 0 {
	// 		continue
	// 	}

	// 	rate, err := strconv.ParseFloat(row[1], 64)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return nil
	// 	}

	// 	record := d.Record{
	// 		DATE:         row[0],
	// 		MORTGAGE30US: rate,
	// 	}

	// 	rows = append(rows, record)

	// }

	return data

}
