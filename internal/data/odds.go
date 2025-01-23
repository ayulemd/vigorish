package data

import (
	"time"
)

type Outcome struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type Market struct {
	Key        string    `json:"key"`
	LastUpdate time.Time `json:"last_update"`
	Outcomes   []Outcome `json:"outcomes"`
}

type Bookmaker struct {
	Key        string    `json:"key"`
	Title      string    `json:"title"`
	LastUpdate time.Time `json:"last_update"`
	Markets    []Market  `json:"markets"`
}

type Odds struct {
	ID         string      `json:"id"`
	SportKey   string      `json:"sport_key"`
	SportTitle string      `json:"sport_title"`
	HomeTeam   string      `json:"home_team"`
	AwayTeam   string      `json:"away_team"`
	Bookmakers []Bookmaker `json:"bookmakers"`
}
