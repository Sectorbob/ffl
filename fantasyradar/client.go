package fantasyradar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client interface {
	GetTournamentSchedule(tournamentID string) (*TournamentSchedule, error)
}

type Config struct {
	APIKey string
}

type httpClient struct {
	Package      string
	AccessLevel  string
	Version      string
	LanguageCode string
	OddsFormat   string
	APIKey       string
	httpClient   *http.Client
}

func (h *httpClient) GetTournamentSchedule(tournamentID string) (*TournamentSchedule, error) {
	baseURL := "https://api.sportradar.com"

	url := fmt.Sprintf("%s/oddscomparison-%s%s%s/%s/%s/tournaments/%s/schedule.json?api_key=%s", baseURL, h.Package, h.AccessLevel, h.Version, h.LanguageCode, h.OddsFormat, tournamentID, h.APIKey)

	res, err := http.Get(url)
	if err != nil {
		panic("unable to send request: " + err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("received error response. code %d", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data TournamentSchedule

	if e := json.Unmarshal(b, &data); e != nil {
		return nil, err
	}

	return &data, nil
}

func NewClient(c *Config) Client {
	return &httpClient{
		Package:      "row", //rest of world (row) united states (us)
		AccessLevel:  "t",   //trial
		Version:      "1",   //version 1
		LanguageCode: "en",
		OddsFormat:   "us",
		APIKey:       c.APIKey,
		httpClient:   http.DefaultClient,
	}
}
