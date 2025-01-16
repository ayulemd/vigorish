package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ayulemd/vigorish/internal/data"
)

var odds []data.Odds

func (app *application) getWithQuery(baseURL string, params map[string]string) error {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return err
	}

	query := parsedURL.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	parsedURL.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")

	res, err := app.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	err = app.readJSON(res, &odds)
	if err != nil {
		return err
	}

	pretty, err := json.MarshalIndent(odds, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(pretty))

	return nil
}
