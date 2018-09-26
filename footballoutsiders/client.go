package footballoutsiders

import (
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Client interface {
	GetWeeklyDVOA(year, week int) (*WeeklyDVOARatings, error)
}

func NewClient() Client {
	return &httpScraper{}
}

type httpScraper struct {
}

func (h *httpScraper) GetWeeklyDVOA(year, week int) (*WeeklyDVOARatings, error) {

	url := fmt.Sprintf("https://www.footballoutsiders.com/dvoa-ratings/%d/week-%d-dvoa-ratings", year, week)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received error code: %d", res.StatusCode)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to parse html doc: %v", err)
	}

	matches := Traverse(doc, atom.Table, []string{"sticky-headers sortable stats", "stats"})

	if len(matches) == 0 {
		return nil, fmt.Errorf("could not find table in html")
	}

	if len(matches) > 1 {
		return nil, fmt.Errorf("matched more than 1 table, unable to determine correct table")
	}

	table := &Table{node: matches[0]}
	data, err := table.Data()
	if err != nil {
		return nil, fmt.Errorf("unable to parse table data: %v", err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("no data rows found in table")
	}

	dvoa := WeeklyDVOARatings{
		Week:    week,
		Year:    year,
		Ratings: make([]DVOARatings, len(data[0])),
	}

	for i, row := range data {
		rank, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, fmt.Errorf("unable to parse value in row[%d] col[0] to int: %v", i, err)
		}

		lastWeekRank, err := strconv.Atoi(row[3])
		if err != nil {
			return nil, fmt.Errorf("unable to parse value in row[%d] col[3] to int: %v", i, err)
		}

		daveRank, err := strconv.Atoi(row[5])
		if err != nil {
			return nil, fmt.Errorf("unable to parse value in row[%d] col[5] to int: %v", i, err)
		}

		offenseRank, err := strconv.Atoi(row[8])
		if err != nil {
			return nil, fmt.Errorf("unable to parse value in row[%d] col[8] to int: %v", i, err)
		}

		defenseRank, err := strconv.Atoi(row[10])
		if err != nil {
			return nil, fmt.Errorf("unable to parse value in row[%d] col[10] to int: %v", i, err)
		}

		stRank, err := strconv.Atoi(row[12])
		if err != nil {
			return nil, fmt.Errorf("unable to parse value in row[%d] col[12] to int: %v", i, err)
		}

		dvoa.Ratings[i] = DVOARatings{
			Rank:             rank,
			TeamAbbrv:        row[1],
			TotalDVOA:        row[2],
			LastWeekRank:     lastWeekRank,
			TotalDAVE:        row[4],
			DAVERank:         daveRank,
			WinLoss:          row[6],
			OffenseDVOA:      row[7],
			OffenseRank:      offenseRank,
			DefenseDVOA:      row[9],
			DefenseRank:      defenseRank,
			SpecialTeamsDVOA: row[11],
			SpecialTeamsRank: stRank,
		}
	}
	return &dvoa, nil
}
