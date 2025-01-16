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

	fmt.Println("Calling API")

	baseURL := "https://api.the-odds-api.com/v4/sports/americanfootball_nfl/odds"

	params := map[string]string{
		"apiKey":     app.config.apiKey,
		"regions":    "us",
		"markets":    "h2h",
		"oddsFormat": "american",
	}

	err = app.getWithQuery(baseURL, params)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
