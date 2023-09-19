package lib

const CsvFormat = ".csv"
const JsonFormat = ".json"

const CountryLevelKey = "Country"
const CampaignLevelKey = "CampaignId"

const PredictedDay = 60.0

const PredictedLinearModel = "linear"

const CountryAggregationLevel = "country"
const CampaignAggregationLevel = "campaign"

type AggregationLevel int
type PredictionModel int

const (
	ByCountry AggregationLevel = iota
	ByCampaign
)

const (
	LinearExtrapolation PredictionModel = iota
)
