package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ayulemd/vigorish/internal/data"
	"github.com/shopspring/decimal"
)

func (app *application) getOdds(baseURL string, params map[string]string) ([]data.Odds, error) {
	var oddsData []data.Odds

	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	query := parsedURL.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	parsedURL.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	res, err := app.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	err = app.readJSON(res.Body, &oddsData)
	if err != nil {
		return nil, err
	}

	return oddsData, nil
}

func (app *application) calculateVig(oddsData []data.Odds) error {
	for _, match := range oddsData {
		fmt.Println(match.AwayTeam, "at", match.HomeTeam)
		fmt.Println("============================================")

		for _, bookmaker := range match.Bookmakers {
			fmt.Println("Bookmaker:", bookmaker.Title)
			fmt.Println("Last Update:", bookmaker.LastUpdate)

			for _, market := range bookmaker.Markets {
				var vig decimal.Decimal

				for _, outcome := range market.Outcomes {
					fmt.Println("Name:", outcome.Name)
					fmt.Println("Price:", outcome.Price)
					impliedProbability := app.impliedProbability(outcome.Price)
					fmt.Printf("Implied Win Probability: %s%%\n", impliedProbability.StringFixed(2))
					vig = vig.Add(impliedProbability)
				}

				vig = vig.Sub(decimal.NewFromInt(100))
				fmt.Printf("The Vigorish: %s%%\n", vig.StringFixed(2))
				fmt.Println("============================================")
			}
		}
	}

	return nil
}
