package transformations

import (
	"fmt"
	d "goetl/data"
	"strconv"
)

func Filter(data [][]string) []d.Record {

	rows := make([]d.Record, 0, len(data))

	//loop through rows and store in map
	for _, row := range data {

		rate, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			fmt.Println(err)
		}

		if rate < 5.0 {
			continue
		} else {

			record := d.Record{
				DATE:         row[0],
				MORTGAGE30US: rate,
			}

			rows = append(rows, record)

			fmt.Println(record.DATE, record.MORTGAGE30US)
		}

	}
	return rows

}

// // Define a function that returns a boolean value
// myCondition := func() bool {
// 	return x > y
// }
