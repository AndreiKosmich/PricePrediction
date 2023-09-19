package services

import (
	"log"
	"priceprediction/lib"
	"strconv"
	"strings"
)

func AggregateData(predictionModel lib.PredictionModel, source string, aggregateLevel lib.AggregationLevel) map[string]float64 {
	if strings.HasSuffix(source, lib.CsvFormat) {
		records := ReadCSVData(source)
		return aggregateCsvData(records, aggregateLevel, predictionModel)
	} else if strings.HasSuffix(source, lib.JsonFormat) {
		records := ReadJSONData(source)
		return aggregateJsonData(records, aggregateLevel, predictionModel)
	} else {
		log.Fatal(lib.UnsupportedFleFormatError)
	}

	return nil
}

func aggregateCsvData(data [][]string, level lib.AggregationLevel, model lib.PredictionModel) map[string]float64 {
	result := make(map[string]float64)
	var levelKey int

	switch level {
	case lib.ByCountry:
		levelKey = 2
	case lib.ByCampaign:
		levelKey = 1
	}

	for _, record := range data {
		country := record[levelKey]
		prices := convertStringsToFloats(record[3:])
		result[country] += predictLTV(prices, model)
	}

	return result
}

func aggregateJsonData(data []map[string]interface{}, level lib.AggregationLevel, model lib.PredictionModel) map[string]float64 {
	result := make(map[string]float64)
	var levelKey string

	switch level {
	case lib.ByCountry:
		levelKey = lib.CountryLevelKey
	case lib.ByCampaign:
		levelKey = lib.CampaignLevelKey
	}

	for _, record := range data {
		country := record[levelKey].(string)
		prices := []float64{
			record["Ltv1"].(float64),
			record["Ltv2"].(float64),
			record["Ltv3"].(float64),
			record["Ltv4"].(float64),
			record["Ltv5"].(float64),
			record["Ltv6"].(float64),
			record["Ltv7"].(float64),
		}

		result[country] += predictLTV(prices, model)
	}

	return result
}

func predictLTV(prices []float64, model lib.PredictionModel) float64 {
	switch model {
	case lib.LinearExtrapolation:
		predictableDays := []float64{1, 2, 3, 4, 5, 6, 7}
		dayIncome := dayIncomePrediction(prices, predictableDays)
		if dayIncome > 0 {
			return dayIncome
		} else {
			return 0.0
		}
	default:
		log.Printf(lib.UnsupportedModelError, model)
		return 0.0
	}
}

func convertStringsToFloats(stringArray []string) []float64 {
	floatArray := make([]float64, len(stringArray))
	for i, str := range stringArray {
		num, err := strconv.ParseFloat(str, 64)

		if err != nil {
			log.Printf(lib.UnsupportedModelError, err)
		}

		floatArray[i] = num
	}
	return floatArray
}

func dayIncomePrediction(x []float64, y []float64) float64 {
	n := len(x)
	sumX, sumY, sumXY, sumX2 := 0.0, 0.0, 0.0, 0.0

	for i := 0; i < n; i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}

	a := (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumX2 - sumX*sumX)
	b := (sumY - a*sumX) / float64(n)

	newX := lib.PredictedDay
	newY := a*newX + b

	return newY
}
