package main

import (
	"fmt"

	"github.com/ayulemd/vigorish/internal/data"
	"github.com/shopspring/decimal"
)

func (app *application) getOdds(baseURL string, params map[string]string) ([]data.Odds, error) {
	var oddsData []data.Odds

	res, err := app.makeApiRequest(baseURL, params)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = app.readJSON(res.Body, &oddsData)
	if err != nil {
		return nil, err
	}

	return oddsData, nil
}

func (app *application) displayOdds(oddsData []data.Odds) error {
	if len(oddsData) == 0 {
		return fmt.Errorf("no odds data provided")
	}

	for _, match := range oddsData {
		fmt.Println(match.AwayTeam, "at", match.HomeTeam)
		fmt.Println("============================================")

		if len(match.Bookmakers) == 0 {
			return fmt.Errorf("no bookmaker data provided")
		}

		for _, bookmaker := range match.Bookmakers {
			fmt.Println("Bookmaker:", bookmaker.Title)
			fmt.Println("Last Update:", bookmaker.LastUpdate)

			if len(bookmaker.Markets) == 0 {
				return fmt.Errorf("no market data provided")
			}

			for _, market := range bookmaker.Markets {
				if len(market.Outcomes) == 0 {
					return fmt.Errorf("no outcomes data provided")
				}

				var impliedProbabilities []decimal.Decimal

				for _, outcome := range market.Outcomes {
					fmt.Println("Name:", outcome.Name)
					fmt.Println("Price:", outcome.Price)
					impliedProbability := app.calculateImpliedProbability(outcome.Price)
					fmt.Printf("Implied Win Probability: %s%%\n", impliedProbability.StringFixed(2))
					impliedProbabilities = append(impliedProbabilities, impliedProbability)
				}

				vig := app.calculateVigorish(impliedProbabilities)
				fmt.Printf("The Vigorish: %s%%\n", vig.StringFixed(2))
				fmt.Println("============================================")
			}
		}
	}

	return nil
}
