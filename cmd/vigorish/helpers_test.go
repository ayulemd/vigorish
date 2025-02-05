package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/shopspring/decimal"
)

var app *application

func TestMain(m *testing.M) {
	fmt.Println("Setting up application instance before tests...")

	app = &application{
		config: config{},
		client: &http.Client{},
	}

	exitCode := m.Run()

	fmt.Println("Cleaning up after tests...")

	os.Exit(exitCode)
}

func TestCalculateImpliedProbability(t *testing.T) {
	testCases := []struct {
		name     string
		price    int64
		expected string
	}{
		{"moneyline", -110, "52.38"},
		{"negative", -250, "71.43"},
		{"positive", 150, "40.00"},
		{"even", 100, "50.00"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := app.calculateImpliedProbability(tc.price)

			expectedDecimal, err := decimal.NewFromString(tc.expected)
			if err != nil {
				t.Fatalf("Invalid expected decimal string %q: %v", tc.expected, err)
			}

			if !result.Equal(expectedDecimal) {
				t.Errorf("impliedProbability(%d) = %s; expected %s", tc.price, result.String(), expectedDecimal.String())
			}
		})
	}
}

func TestCalculateVigorish(t *testing.T) {
	testCases := []struct {
		name                 string
		impliedProbabilities []decimal.Decimal
		expected             string
	}{
		{
			"standard -110 moneyline",
			[]decimal.Decimal{
				decimal.RequireFromString("52.38"),
				decimal.RequireFromString("52.38"),
			},
			"4.76",
		},
		{
			"-124 vs. 106 moneyline",
			[]decimal.Decimal{
				decimal.RequireFromString("55.36"),
				decimal.RequireFromString("48.54"),
			},
			"3.90",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := app.calculateVigorish(tc.impliedProbabilities)

			expectedDecimal, err := decimal.NewFromString(tc.expected)
			if err != nil {
				t.Fatalf("Invalid expected decimal string %q: %v", tc.expected, err)
			}

			if !result.Equal(expectedDecimal) {
				t.Errorf("Implied probabilities: %v,  vigorish = %s; expected %s", tc.impliedProbabilities, result.String(), expectedDecimal.String())
			}
		})
	}
}
