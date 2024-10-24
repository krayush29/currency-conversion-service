package services

import (
	"math"
)

type CurrencyType string

const (
	USD CurrencyType = "USD"
	INR CurrencyType = "INR"
)

var conversionRates = map[CurrencyType]float64{
	USD: 1.0,
	INR: 85.0,
}

type ConversionRequest struct {
	FromCurrency CurrencyType `json:"from_currency"`
	ToCurrency   CurrencyType `json:"to_currency"`
	Amount       float64      `json:"amount"`
}

type ConversionResponse struct {
	ConvertedAmount float64 `json:"converted_amount"`
}

func roundOf(value float64) float64 {
	return math.Round(value*100) / 100
}

func Convert(fromCurrency, toCurrency CurrencyType, amount float64) float64 {
	baseValue := amount / conversionRates[fromCurrency]
	return roundOf(baseValue * conversionRates[toCurrency])
}
