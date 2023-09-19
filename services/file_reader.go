package services

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"priceprediction/lib"
)

func ReadCSVData(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf(lib.CsvOpeningError, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf(lib.CsvOpeningError, err)
	}

	return records[1:]
}

func ReadJSONData(filename string) []map[string]interface{} {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf(lib.JsonOpeningError, err)
	}

	var records []map[string]interface{}
	err = json.Unmarshal(file, &records)
	if err != nil {
		log.Fatalf(lib.JsonOpeningError, err)
	}

	return records
}
