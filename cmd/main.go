package main

import (
	"fmt"
	src "goetl/pkg/sources"
	tgt "goetl/pkg/targets"
	tr "goetl/pkg/transformations"
)

func main() {

	fmt.Println("Hi go-etl!\nEasy pipelines to move data from source to target!")

	data := src.CsvSrc("data/MORTGAGE30US.csv")

	//Filter condiditon is only changed in the function currently
	//working on way to pass conditions from main
	highRateData := tr.Filter(data)

	tgt.CsvTgt("HIGHRATEDATA", highRateData)

}
