package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// 1. Create a new struct for storing read JSON objects
type SampleJSONFormat struct {
	Sample1 string `json:"Sample1"`
	Sample2 string `json:"Sample2"`
	Sample3 int64  `json:"Sample3"`
}

func convertJSONToCSV(source, destination string) error {
	// 2. Read the JSON file into the struct array
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	var SampleJSON []SampleJSONFormat
	if err := json.NewDecoder(sourceFile).Decode(&SampleJSON); err != nil {
		return err
	}

	// 3. Create a new CSV file
	// outputFile, err := os.Create(destination)
	// if err != nil {
	// 	return err
	// }
	// defer outputFile.Close()

	// 4. Create a header for the new CSV file
	// writer := csv.NewWriter(outputFile)
	// defer writer.Flush()

	// header := []string{"Sample1", "Sample2", "Sample3"}
	// if err := writer.Write(header); err != nil {
	// 	return err
	// }

	// 3. Open the existing CSV file
	//var path = "from_json.csv"
	var path = destination
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	// 4. Append JSON data into the CSV file
	writer := csv.NewWriter(f)
	defer writer.Flush()

	for _, r := range SampleJSON {
		var csvRow []string
		csvRow = append(csvRow, r.Sample1, r.Sample2, fmt.Sprint(r.Sample3))
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	start := time.Now()
	files, err := ioutil.ReadDir("./data/")
	for _, file := range files {
		if err := convertJSONToCSV("./data/"+file.Name(), "from_json.csv"); err != nil {
			log.Fatal(err)
		}
	}
	elapsed := time.Since(start)
	log.Printf("%s", elapsed)
	if err != nil {
		log.Fatal(err)
	}
}
