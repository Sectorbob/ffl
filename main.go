package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/Sectorbob/ffl/fantasyradar"
	"github.com/Sectorbob/ffl/footballoutsiders"
	sf "github.com/joelhill/mysportsfeeds-go"
)

const (
	nflTournmentID = "sr:tournament:31"
)

var sportsRadarAPIKey = os.Getenv("SPORTSRADAR_API_KEY")
var mySportsFeedAPIKey = os.Getenv("MSF_API_KEY")
var mySportsFeedPassword = os.Getenv("MSF_PASSWORD")

func main() {

	client := fantasyradar.NewClient(&fantasyradar.Config{APIKey: sportsRadarAPIKey})

	data, err := client.GetTournamentSchedule(nflTournmentID)
	if err != nil {
		panic(err)
	}

	for _, v := range data.SportEvents {
		fmt.Println(v.String())
	}

	config := &sf.Config{
		Authorization: auth(mySportsFeedAPIKey, mySportsFeedPassword),
		BaseURL:       sf.MySportsFeedBaseURL,
		Version:       sf.VersionV1_2,
		Sport:         sf.SportNFL,
		Format:        sf.FormatJSON,
		Season:        sf.SeasonCurrent,
	}

	sfClient := sf.NewService(config)

	c := context.Background()
	gameOpts := sfClient.NewSeasonalGamesOptions()

	games, code, err := sfClient.SeasonalGames(c, gameOpts)
	fmt.Printf("code: %d\n", code)
	fmt.Printf("err:  %v\n", err)
	for _, v := range *games.Games {
		fmt.Printf("%s @ %s\n", v.Schedule.AwayTeam.Abbreviation, v.Schedule.HomeTeam.Abbreviation)
	}

	foClient := footballoutsiders.NewClient()

	dvoa, err := foClient.GetWeeklyDVOA(2018, 1)
	if err != nil {
		panic("unable to pull dvoa: " + err.Error())
	}

	for _, v := range dvoa.Ratings {
		fmt.Printf("Team: %s, DVOA: %s\n", v.TeamAbbrv, v.TotalDVOA)
	}
}

func auth(user, pass string) string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+pass))
}
