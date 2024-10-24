package services

import (
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name           string
		fromCurrency   CurrencyType
		toCurrency     CurrencyType
		amount         float64
		expectedAmount float64
	}{
		{
			name:           "USD to INR",
			fromCurrency:   USD,
			toCurrency:     INR,
			amount:         100,
			expectedAmount: 8500, // 100 USD * 85 INR/USD
		},
		{
			name:           "INR to USD",
			fromCurrency:   INR,
			toCurrency:     USD,
			amount:         8500,
			expectedAmount: 100, // 8500 INR / 85 INR/USD
		},
		{
			name:           "USD to USD",
			fromCurrency:   USD,
			toCurrency:     USD,
			amount:         100,
			expectedAmount: 100, // 100 USD * 1 USD/USD
		},
		{
			name:           "INR to INR",
			fromCurrency:   INR,
			toCurrency:     INR,
			amount:         8500,
			expectedAmount: 8500, // 8500 INR * 1 INR/INR
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Convert(tt.fromCurrency, tt.toCurrency, tt.amount)
			if result != tt.expectedAmount {
				t.Errorf("Convert(%v, %v, %v) = %v; want %v", tt.fromCurrency, tt.toCurrency, tt.amount, result, tt.expectedAmount)
			}
		})
	}
}

func TestRoundOf(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected float64
	}{
		{
			name:     "Round down",
			value:    123.456,
			expected: 123.46,
		},
		{
			name:     "Round up",
			value:    123.454,
			expected: 123.45,
		},
		{
			name:     "No rounding needed",
			value:    123.450,
			expected: 123.45,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := roundOf(tt.value)
			if result != tt.expected {
				t.Errorf("roundOf(%v) = %v; want %v", tt.value, result, tt.expected)
			}
		})
	}
}
