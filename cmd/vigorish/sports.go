package main

import (
	"fmt"

	"github.com/ayulemd/vigorish/internal/data"
)

func (app *application) getSports(baseURL string, params map[string]string) ([]data.Sport, error) {
	var sports []data.Sport

	res, err := app.makeApiRequest(baseURL, params)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = app.readJSON(res.Body, &sports)
	if err != nil {
		return nil, err
	}

	return sports, nil
}

func (app *application) displaySports(sports []data.Sport) {
	for i, sport := range sports {
		fmt.Printf("%d. %s - %s\n", i+1, sport.Group, sport.Title)
	}
}

func (app *application) selectSport(sports []data.Sport) (string, error) {
	var sportIndex int

	fmt.Print("Select sport: ")
	_, err := fmt.Scanln(&sportIndex)
	if err != nil {
		return "", err
	}

	return sports[sportIndex-1].Key, nil
}
