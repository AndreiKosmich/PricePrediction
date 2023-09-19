package modules

import (
	"flag"
	"fmt"
	"log"
	"priceprediction/lib"
)

func HandleParams() (lib.PredictionModel, string, lib.AggregationLevel) {
	var modelStr, source, aggregateStr string
	flag.StringVar(&modelStr, "model", "", "Prediction model")
	flag.StringVar(&source, "source", "", "File path")
	flag.StringVar(&aggregateStr, "aggregate", "", "Data aggregation level")
	flag.Parse()

	model, err := parseModel(modelStr)
	if err != nil {
		log.Fatalf(lib.ParsingError, err)
	}

	aggregateLevel, err := parseAggregationLevel(aggregateStr)
	if err != nil {
		log.Fatalf(lib.AggregationError, err)
	}

	return model, source, aggregateLevel
}

func parseModel(modelStr string) (lib.PredictionModel, error) {
	switch modelStr {
	case lib.PredictedLinearModel:
		return lib.LinearExtrapolation, nil
	default:
		return -1, fmt.Errorf(lib.UnsupportedModelError, modelStr)
	}
}

func parseAggregationLevel(aggregateStr string) (lib.AggregationLevel, error) {
	switch aggregateStr {
	case lib.CountryAggregationLevel:
		return lib.ByCountry, nil
	case lib.CampaignAggregationLevel:
		return lib.ByCampaign, nil
	default:
		return -1, fmt.Errorf(lib.AggregationError, aggregateStr)
	}
}
