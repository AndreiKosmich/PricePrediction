# Forecasting LTV (Life Time Value) of the 60th day

This is a utility written in Go that allows you to predict the expected day 60 profit per user based on data from various sources. The utility supports two different prediction methods and can work with data from CSV and JSON files. You can run the utility via `go run` without additional manipulations.

## Data description

There are two sources of user data:

1. **CSV file**: Contains the following fields - user id, advertising campaign id, country, as well as the sequence of profit (LTV - life time value) by day that the user brought. For example, Ltv7 means that the user brought us x dollars by the 7th day.

2. **JSON file**: Contains similar data, but already averaged. That is, advertising campaign id, country, average LTV of users and their number. Ltv7 here also means that on average users from that country and campaigns brought us x dollars by day 7. The total profit of the 7th day is calculated as x multiplied by the number of users.

Both data sources are independent, so results may vary.

## Task

It is necessary to write a utility that predicts day 60 LTV based on data from the specified sources. The utility must support at least two different prediction methods. The accuracy of the prediction does not matter, but the only condition is that the prediction result must be non-decreasing.

You can use simple linear extrapolation or third party libraries if needed.

## Usage

The utility supports the following command line switches:

- `-model`: Select one of the prediction methods (for example, "linear").
- `-source`: Name of the data file (CSV or JSON).
- `-aggregate`: Select the data aggregation level (country or campaign).

Example of running the utility:

```sh
go run main.go -model linear -source data.csv -aggregate country
