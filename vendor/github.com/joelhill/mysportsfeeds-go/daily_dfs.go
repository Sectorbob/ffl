package msf

import (
	"context"
	"errors"
	"fmt"
	"time"

	blaster "github.com/joelhill/go-rest-http-blaster"
	logrus "github.com/sirupsen/logrus"
)

// DailyDfsOptions - Are the options to hit the daily DFS endpoint
type DailyDfsOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Season  string
	Date    string
	Format  string

	// Optional URL Params
	Team     string
	Player   string
	Position string
	Country  string
	DfsType  string
	Sort     string
	Offset   string
	Limit    string
	Force    string
}

// NewDailyDfsOptions - Returns the options with most url parts already set to hit the daily DFS endpoint
func (s *Service) NewDailyDfsOptions() *DailyDfsOptions {
	return &DailyDfsOptions{
		URL:     s.Config.BaseURL,
		Version: s.Config.Version,
		Sport:   s.Config.Sport,
		Format:  s.Config.Format,
		Season:  s.Config.Season,
	}
}

// DailyDfs - hits the https://api.mysportsfeeds.com/{version/pull/{sport}/{season}/date/{date}/dfs.{format} endpoint
func (s *Service) DailyDfs(c context.Context, options *DailyDfsOptions) (DfsIO, int, error) {
	errorPayload := make(map[string]interface{})
	mapping := DfsIO{}

	// make sure we have all the required elements to build the full required url string.
	err := validateDailyDfsURI(options)
	if err != nil {
		return mapping, 0, err
	}

	t := time.Now()
	cacheBuster := t.Format("20060102150405")

	uri := fmt.Sprintf("%s/%s/pull/%s/%s/date/%s/dfs.%s?cachebuster=%s", options.URL, options.Version, options.Sport, options.Season, options.Date, options.Format, cacheBuster)

	if len(options.Team) > 0 {
		uri = fmt.Sprintf("%s&team=%s", uri, options.Team)
	}

	if len(options.Player) > 0 {
		uri = fmt.Sprintf("%s&player=%s", uri, options.Player)
	}

	if len(options.Position) > 0 {
		uri = fmt.Sprintf("%s&position=%s", uri, options.Position)
	}

	if len(options.Country) > 0 {
		uri = fmt.Sprintf("%s&country=%s", uri, options.Country)
	}

	if len(options.DfsType) > 0 {
		uri = fmt.Sprintf("%s&dfstype=%s", uri, options.DfsType)
	}

	if len(options.Sort) > 0 {
		uri = fmt.Sprintf("%s&sort=%s", uri, options.Sort)
	}

	if len(options.Offset) > 0 {
		uri = fmt.Sprintf("%s&offset=%s", uri, options.Offset)
	}

	if len(options.Limit) > 0 {
		uri = fmt.Sprintf("%s&limit=%s", uri, options.Limit)
	}

	if len(options.Force) > 0 {
		uri = fmt.Sprintf("%s&force=%s", uri, options.Force)
	}

	s.Logger = s.Logger.WithFields(logrus.Fields{
		"URI": uri,
	})
	s.Logger.Debug("Seasonal DFS API Call")

	// make you a client
	client, err := blaster.NewClient(uri)
	if err != nil {
		s.Logger.Errorf("failed to create a http client: %s", err.Error())
		return mapping, 0, err
	}

	client.SetHeader("Accept-Encoding", CompressionHeaderGzip)
	client.SetHeader("Authorization", s.Config.Authorization)
	client.WillSaturateOnError(&errorPayload)
	client.WillSaturate(&mapping)

	ctx := context.Background()
	statusCode, err := client.Get(ctx)
	if err != nil {
		s.Logger.Errorf("something went wrong making the get request for DailyDfs: %s", err.Error())
		return mapping, statusCode, err
	}

	s.Logger.Infof("DailyDfs Status Code: %d", statusCode)

	if client.StatusCodeIsError() {
		s.Logger.Errorf("DailyDfs retuned an unsuccessful status code. Error: %+v", errorPayload)
	}

	return mapping, statusCode, nil
}

func validateDailyDfsURI(options *DailyDfsOptions) error {
	if len(options.URL) == 0 {
		return errors.New("missing required option to build the url: URL")
	}
	if len(options.Version) == 0 {
		return errors.New("missing required option to build the url: Version")
	}
	if len(options.Sport) == 0 {
		return errors.New("missing required option to build the url: Sport")
	}
	if len(options.Season) == 0 {
		return errors.New("missing required option to build the url: Season")
	}
	if len(options.Date) == 0 {
		return errors.New("missing required option to build the url: Date")
	}
	if len(options.Format) == 0 {
		return errors.New("missing required option to build the url: Format")
	}
	return nil
}
