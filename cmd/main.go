package main

import (
	"fmt"
	src "goetl/pkg/sources"
)

func main() {

	fmt.Println("Hi go-etl!\nEasy pipelines to move data from source to target!")

	src.CsvSrc("data/MORTGAGE30US.csv")
}
