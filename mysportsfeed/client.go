package mysportsfeed

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

const V1_2 = "v1.2"

type Client interface {
	GetBoxScore(gameID string) (*GameBoxScore, error)
}

type Config struct {
	BaseURL    string
	APIToken   string
	Password   string
	Version    string
	Sport      string
	Tournament string
}

type httpClient struct {
	Conf   Config
	client *http.Client
}

func (h *httpClient) GetBoxScore(gameID string) (*GameBoxScore, error) {

	url := fmt.Sprintf("%s/%s/pull/%s/%s/game_boxscore.json?gameid=%s", h.Conf.BaseURL, h.Conf.Version, h.Conf.Sport, h.Conf.Tournament, gameID)

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

}
