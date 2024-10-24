package controllers

import (
	"currency-conversion-service/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	var req services.ConversionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.FromCurrency == "" || req.ToCurrency == "" || req.Amount == 0 {
		http.Error(w, "Missing required fields: from_currency, to_currency, amount", http.StatusBadRequest)
		return
	}

	// Validate currency types
	if !isValidCurrencyType(req.FromCurrency) || !isValidCurrencyType(req.ToCurrency) {
		validCurrencies := []services.CurrencyType{services.USD, services.INR}
		http.Error(w, fmt.Sprintf("Invalid currency type: %s or %s. Please provide a valid currency type. Valid types are: %v", req.FromCurrency, req.ToCurrency, validCurrencies), http.StatusBadRequest)
		return
	}

	convertedAmount := services.Convert(req.FromCurrency, req.ToCurrency, req.Amount)
	resp := services.ConversionResponse{ConvertedAmount: convertedAmount}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func isValidCurrencyType(currency services.CurrencyType) bool {
	switch currency {
	case services.USD, services.INR:
		return true
	default:
		return false
	}
}
