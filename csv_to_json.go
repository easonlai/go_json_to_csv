package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type SampleJSONFormat struct {
	Sample1 string
	Sample2 string
	Sample3 int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	csvFile, err := os.Open("./to_json.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var record SampleJSONFormat
	var records []SampleJSONFormat

	for _, each := range csvData {
		record.Sample1 = each[0]
		record.Sample2 = each[1]
		record.Sample3, _ = strconv.Atoi(each[2])
		records = append(records, record)

		jsonData, err := json.Marshal(records)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		jsonFile, err := os.Create("./data/data" + strconv.Itoa(rand.Int()) + ".json")
		//jsonFile, err := os.Create("./data/data" + strconv.Itoa(rand.Intn(100000000)) + ".json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		jsonFile.Write(jsonData)
		jsonFile.Close()
		records = nil
	}
}
