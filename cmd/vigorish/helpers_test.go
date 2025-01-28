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
	app = &application{
		config: config{},
		client: &http.Client{},
	}

	fmt.Println("Setting up application instance before tests...")

	exitCode := m.Run()

	fmt.Println("Cleaning up after tests...")

	os.Exit(exitCode)
}

func TestImpliedProbability(t *testing.T) {
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
			result := app.impliedProbability(tc.price)

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
