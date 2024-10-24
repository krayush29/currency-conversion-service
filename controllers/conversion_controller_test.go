package controllers

import (
	"bytes"
	"currency-conversion-service/services"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConvertHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    services.ConversionRequest
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Valid request",
			requestBody: services.ConversionRequest{
				FromCurrency: "USD",
				ToCurrency:   "INR",
				Amount:       100,
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "{\"converted_amount\":8500}\n", // Assuming conversion rate is 1 USD = 85 INR
		},
		{
			name: "Missing required fields",
			requestBody: services.ConversionRequest{
				FromCurrency: "",
				ToCurrency:   "INR",
				Amount:       100,
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Missing required fields: from_currency, to_currency, amount\n",
		},
		{
			name: "Invalid currency type",
			requestBody: services.ConversionRequest{
				FromCurrency: "EUR",
				ToCurrency:   "INR",
				Amount:       100,
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid currency type: EUR or INR. Please provide a valid currency type. Valid types are: [USD INR]\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, _ := json.Marshal(test.requestBody)
			req, err := http.NewRequest("POST", "/convert", bytes.NewBuffer(body))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(ConvertHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != test.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, test.expectedStatus)
			}

			if rr.Body.String() != test.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), test.expectedBody)
			}
		})
	}
}
