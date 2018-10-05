package yahoo

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const DefaultURL = "https://dfyql-ro.sports.yahoo.com/v2/export/contestPlayers"

type Client interface {
	GetPlayers(contestID int) ([]Player, error)
}

func NewClient() Client {
	return &httpClient{
		BaseURL: DefaultURL,
		client:  http.DefaultClient,
	}
}

type httpClient struct {
	BaseURL string
	client  *http.Client
}

func (h *httpClient) GetPlayers(contestID int) ([]Player, error) {
	// Build out url
	path, err := url.Parse(h.BaseURL)
	if err != nil {
		panic(err)
	}
	q := path.Query()
	q.Add("contestId", strconv.Itoa(contestID))
	path.RawQuery = q.Encode()

	res, err := h.client.Get(path.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got status code %d", res.StatusCode)
	}

	// Parse CSV

	// Read File into a Variable
	lines, err := csv.NewReader(res.Body).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to parse csv response: %v", err)
	}

	// Loop through lines & turn into object
	players := make([]Player, len(lines)-1)
	for i, line := range lines {
		if i == 0 {
			continue // ignore header row
		}

		if len(line) != 12 {
			return nil, fmt.Errorf("csv row %d did not have 12 columns. actual: %d", i, len(line))
		}

		salary, err := strconv.Atoi(line[8])
		if err != nil {
			return nil, fmt.Errorf("row %d, col %d was unable to be parsed into an integer: %v", i, 8, err)
		}
		fppg, err := strconv.ParseFloat(line[9], 32)
		if err != nil {
			return nil, fmt.Errorf("row %d, col %d was unable to be parsed into a float: %v", i, 9, err)
		}
		starting := strings.ToLower(line[11]) == "yes"

		players[i-1] = Player{
			ID:           line[0],
			FirstName:    line[1],
			LastName:     line[2],
			Position:     line[3],
			Team:         line[4],
			Opponent:     line[5],
			Game:         line[6],
			Time:         line[7],
			Salary:       salary,        // line[8]
			FPPG:         float32(fppg), // line[9]
			InjuryStatus: line[10],
			Starting:     starting, // line[11]
		}
	}

	return players, nil
}

// Player is the daily fantasy sports player info including the price.
type Player struct {
	ID           string
	FirstName    string
	LastName     string
	Position     string
	Team         string
	Opponent     string
	Game         string
	Time         string
	Salary       int
	FPPG         float32
	InjuryStatus string
	Starting     bool
}

// FullName formats the full player name
func (p Player) FullName() string {
	return strings.TrimSpace(p.FirstName + " " + p.LastName)
}

// IsNFL checks if the player's id denotes an nfl player
func (p Player) IsNFL() bool {
	parts := strings.Split(p.ID, ".")
	if len(parts) == 0 {
		return false
	}
	return parts[0] == "nfl"
}

func (p Player) HomeTeam() string {
	parts := strings.Split(p.Game, "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

func (p Player) AwayTeam() string {
	parts := strings.Split(p.Game, "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[0]
}
