package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	apiKey string
}

type application struct {
	config config
	client *http.Client
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg config

	cfg.apiKey = os.Getenv("THE_ODDS_API_KEY")

	app := &application{
		config: cfg,
		client: &http.Client{},
	}

	baseURL := "https://api.the-odds-api.com/v4/sports"

	params := map[string]string{
		"apiKey":     app.config.apiKey,
		"regions":    "us",
		"oddsFormat": "american",
	}

	sports, err := app.getSports(baseURL, params)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	app.displaySports(sports)

	key, err := app.selectSport(sports)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	baseURL = fmt.Sprintf("https://api.the-odds-api.com/v4/sports/%s/odds", key)

	oddsData, err := app.getOdds(baseURL, params)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	app.displayOdds(oddsData)
	if err != nil {
		fmt.Print("Error:", err)
		os.Exit(1)
	}
}
