package fantasyradar

import (
	"fmt"
	"time"
)

var newyork *time.Location

func init() {
	var err error
	newyork, err = time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
}

// TournamentSchedule TODO: write me
type TournamentSchedule struct {
	GeneratedAt string `json:"generated_at"`
	Schema      string `json:"schema"`
	Tournament
}

// Tournament TODO: write me
type Tournament struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Sport       Sport        `json:"sport"`
	SportEvents []SportEvent `json:"sport_events"`
}

// Sport TODO: write me
type Sport struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// SportEvent TODO: write me
type SportEvent struct {
	ID                 string          `json:"id"`
	Scheduled          string          `json:"scheduled"`
	StartTimeTBD       bool            `json:"start_time_tbd"`
	Status             string          `json:"status"`
	TournamentRound    TournamentRound `json:"tournament_round"`
	Season             Season          `json:"season"`
	Tournment          Tournament      `json:"tournament"`
	Competitors        []Competitor    `json:"competitors"`
	Venue              Venue           `json:"venue"`
	UUIDs              string          `json:"uuids"`
	Markets            []Market        `json:"market"`
	MarketsLastUpdated string          `json:"markets_last_updated"`
	Consensus          Consensus       `json:"consensus"`
}

func (s SportEvent) String() string {
	date := s.Start().In(newyork).Format("Mon 01/02/2006")
	time := s.Start().In(newyork).Format("03:04:05pm")

	return fmt.Sprintf("Week: %d Date: %s Time: %s Home: %-20s Away: %-20s Spread: %5s O/U: %4s", s.TournamentRound.Number, date, time, s.Home().Name, s.Away().Name, s.CurrentSpread(), s.OverUnder())
}

func (s SportEvent) Home() *Competitor {
	for _, v := range s.Competitors {
		if v.Qualifier == "home" {
			return &v
		}
	}
	return nil
}

func (s SportEvent) Away() *Competitor {
	for _, v := range s.Competitors {
		if v.Qualifier == "away" {
			return &v
		}
	}
	return nil
}

func (s SportEvent) CurrentSpread() string {
	for _, v := range s.Consensus.Lines {
		if v.Name == "spread_current" {
			return v.Spread
		}
	}
	return "N/A"
}

func (s SportEvent) OverUnder() string {
	for _, v := range s.Consensus.Lines {
		if v.Name == "total_current" {
			return v.Outcomes[0].Total
		}
	}
	return "N/A"
}

func (s SportEvent) Start() time.Time {
	// 2018-09-25T00:15:00+00:00

	t, err := time.Parse("2006-01-02T15:04:05-07:00", s.Scheduled)
	if err != nil {
		panic(err)
	}

	return t
}

// TournamentRound TODO: write me
type TournamentRound struct {
	Type   string
	Number int // week
}

// Season TODO: write me
type Season struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Year         string `json:"year"`
	TournamentID string `json:"tournament_id"`
	UUIDs        string `json:"uuids"`
}

// Competitor TODO: write me
type Competitor struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Country        string `json:"country"`
	CountryCode    string `json:"country_code"`
	Abbreviation   string `json:"abbreviation"`
	Qualifier      string `json:"qualifier"`
	RotationNumber int    `json:"rotation_number"`
	UUIDs          string `json:"uuids"`
}

// Venue TODO: write me
type Venue struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Capacity       int    `json:"capacity"`
	CityName       string `json:"city_name"`
	CountryName    string `json:"country_name"`
	MapCoordinates string `json:"map_coordinates"`
	CountryCode    string `json:"country_code"`
	UUIDs          string `json:"uuids"`
}

type Market struct {
	//TODO: define
}

type Consensus struct {
	BetPercentageOutcomes []Outcome `json:"bet_percentage_outcomes"`
	Lines                 []Outcome `json:"lines"`
}

type Outcome struct {
	Name     string         `json:"name"`
	Spread   string         `json:"spread"`
	Outcomes []OutcomeEntry `json:"outcomes"`
}

type OutcomeEntry struct {
	Type       string `json:"type"`
	Percentage int    `json:"percentage"`
	Odds       string `json:"odds"`
	Total      string `json:"total"`
}
