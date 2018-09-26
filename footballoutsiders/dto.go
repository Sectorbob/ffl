package footballoutsiders

type DVOARatings struct {
	Rank             int
	TeamAbbrv        string
	TotalDVOA        string
	LastWeekRank     int
	TotalDAVE        string
	DAVERank         int
	WinLoss          string
	OffenseDVOA      string
	OffenseRank      int
	DefenseDVOA      string
	DefenseRank      int
	SpecialTeamsDVOA string
	SpecialTeamsRank int
}

type WeeklyDVOARatings struct {
	Year    int
	Week    int
	Ratings []DVOARatings
}
