package mysportsfeed

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

const V1_2 = "v1.2"

type Client interface {
	GetBoxScore(season, gameID string) (*GameBoxScore, error)
	CumulativePlayersStats(season string) (*CumulativePlayerStats, error)
	GetSchedule(season string) (*FullGameSchedule, error)
}

type Config struct {
	BaseURL  string
	APIToken string
	Password string
	Version  string
	Sport    string
}

func NewClient(conf Config) Client {
	return &httpClient{
		Conf:   conf,
		client: http.DefaultClient,
	}
}

type httpClient struct {
	Conf   Config
	client *http.Client
}

func (h *httpClient) GetBoxScore(season, gameID string) (*GameBoxScore, error) {

	url := fmt.Sprintf("%s/%s/pull/%s/%s/game_boxscore.json?gameid=%s", h.Conf.BaseURL, h.Conf.Version, h.Conf.Sport, season, gameID)

	// Create Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Attach auth header
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(h.Conf.APIToken+":"+h.Conf.Password)))

	// Send Request
	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error code: " + res.Status) //TODO: fix
	}

	//TODO: write the rest

	var boxScore GameBoxScoreWrapper
	if e := json.NewDecoder(res.Body).Decode(&boxScore); e != nil {
		return nil, e
	}

	return &boxScore.GameBoxScore, nil
}

func (h *httpClient) CumulativePlayersStats(season string) (*CumulativePlayerStats, error) {
	url := fmt.Sprintf("%s/%s/pull/%s/%s/cumulative_player_stats.json", h.Conf.BaseURL, h.Conf.Version, h.Conf.Sport, season)

	// Create Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Attach auth header
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(h.Conf.APIToken+":"+h.Conf.Password)))

	// Send Request
	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error code: " + res.Status) //TODO: fix
	}

	//TODO: write the rest

	var cpsw CumulativePlayerStatsWrapper
	if e := json.NewDecoder(res.Body).Decode(&cpsw); e != nil {
		return nil, e
	}

	return &cpsw.CumulativePlayerStats, nil
}

func (h *httpClient) GetSchedule(season string) (*FullGameSchedule, error) {
	url := fmt.Sprintf("%s/%s/pull/%s/%s/full_game_schedule.json", h.Conf.BaseURL, h.Conf.Version, h.Conf.Sport, season)

	// Create Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Attach auth header
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(h.Conf.APIToken+":"+h.Conf.Password)))

	// Send Request
	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error code: " + res.Status) //TODO: fix
	}

	//TODO: write the rest

	var fgsw FullGameScheduleWrapper
	if e := json.NewDecoder(res.Body).Decode(&fgsw); e != nil {
		return nil, e
	}

	return &fgsw.FullGameSchedule, nil
}
