package main

import (
	"fmt"
	"priceprediction/lib"
	"priceprediction/modules"
	"priceprediction/services"
)

func main() {
	predictionModel, source, aggregateLevel := modules.HandleParams()
	aggregatedData := services.AggregateData(predictionModel, source, aggregateLevel)
	resultFormat := resultFormat(aggregateLevel)

	for key, value := range aggregatedData {
		fmt.Printf(resultFormat, key, value)
	}
}

func resultFormat(aggregateLevel lib.AggregationLevel) string {
	switch aggregateLevel {
	case lib.ByCountry:
		return "%s: %.2f\n"
	case lib.ByCampaign:
		return "<campaign %s>: %.2f\n"
	default:
		return "Unknown format\n" // Обработка случая, когда уровень агрегации неизвестен
	}
}
