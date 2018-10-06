package mysportsfeed

import (
	"fmt"
)

type FullGameSchedule struct {
	LastUpdatedOn string      `json:"lastUpdatedOn"`
	GameEntries   []GameEntry `json:"gameentry"`
}

type FullGameScheduleWrapper struct {
	FullGameSchedule FullGameSchedule `json:"fullgameschedule"`
}

type GameEntry struct {
	ID                       string `json:"id"`
	Week                     string `json:"week"`
	ScheduleStatus           string `json:"scheduleStatus"`
	OriginalDate             string `json:"originalDate"`
	OriginalTime             string `json:"originalTime"`
	DelayedOrPostponedReason string `json:"delayedOrPostponedReason"`
	Date                     string `json:"date"`
	Time                     string `json:"time"`
	AwayTeam                 Team   `json:"awayTeam"`
	HomeTeam                 Team   `json:"homeTeam"`
	Location                 string `json:"location"`
}

type Team struct {
	ID           string `json:"ID"`
	City         string `json:"City"`
	Name         string `json:"Name"`
	Abbreviation string `json:"Abbreviation"`
}

type GameBoxScoreWrapper struct {
	GameBoxScore GameBoxScore `json:"gameboxscore"`
}

type GameBoxScore struct {
	LastUpdatedOn  string          `json:"lastUpdatedOn"`
	Game           GameEntry       `json:"game"`
	QuarterSummary QuarterSummary  `json:"quarterSummary"`
	AwayTeam       AwayTeamSummary `json:"awayTeam"`
	HomeTeam       HomeTeamSummary `json:"homeTeam"`
}

type QuarterSummary struct {
	Quarters []Quarter     `json:"quarter"`
	Totals   QuarterTotals `json:"quarterTotals"`
}

type Quarter struct {
	Number       string        `json:"@number"`
	AwayScore    string        `json:"awayScore"`
	HomeScore    string        `json:"homeScore"`
	ScoringPlays []ScoringPlay `json:"scoringPlay"`
}

type QuarterTotals struct {
	AwayScore string `json:"awayScore"`
	HomeScore string `json:"homeScore"`
}

type ScoringPlay struct {
	Time             string `json:"time"`
	TeamAbbreviation string `json:"teamAbbreviation"`
	PlayDescription  string `json:"playDescription"`
}

type AwayTeamSummary struct {
	Stats   TeamStats          `json:"awayTeamStats"`
	Players PlayerEntryWrapper `json:"awayPlayers"`
}

type HomeTeamSummary struct {
	Stats   TeamStats          `json:"homeTeamStats"`
	Players PlayerEntryWrapper `json:"homePlayers"`
}

// TeamStats TODO: write me
type TeamStats struct {
	PassingStats
	QBR Stat `json:"QBRating"`
	RushingStats
	ReceivingStats
	DefenseStats
	KB Stat `json:"KB"`
	FumbleStats
	KickoffReturnStats
	PuntReturnStats
	KickingStats
	KickoffStats
	PuntStats
	TeamCumulatives
	PointsFor         Stat `json:"PointsFor"`
	PointsAgainst     Stat `json:"PointsAgainst"`
	PointDifferential Stat `json:"PointDifferential"`
	TwoPtStats
}

type PlayerStats struct {
	GamesPlayed Stat `json:"GamesPlayed"`
	PassingStats
	QBR Stat `json:"QBRating"`
	RushingStats
	ReceivingStats
	DefenseStats
	FumbleStats
	KickoffReturnStats
	PuntReturnStats
	TwoPtStats
}

func (p PlayerStats) String() string {
	return fmt.Sprintf("G: %2s, Passing:[Att:%3s Cmp:%3s]", p.GamesPlayed.Text, p.PassAttempts.Text, p.PassCompletions)
}

type TwoPtStats struct {
	TwoPtAtt      Stat `json:"TwoPtAtt"`
	TwoPtMade     Stat `json:"TwoPtMade"`
	TwoPtPassAtt  Stat `json:"TwoPtPassAtt"`
	TwoPtPassMade Stat `json:"TwoPtPassMade"`
	TwoPtRushAtt  Stat `json:"TwoPtRushAtt"`
	TwoPtRushMade Stat `json:"TwoPtRushMade"`
}

type TeamCumulatives struct {
	FirstDownsTotal   Stat `json:"FirstDownsTotal"`
	FirstDownsPass    Stat `json:"FirstDownsPass"`
	FirstDownsRush    Stat `json:"FirstDownsRush"`
	FirstDownsPenalty Stat `json:"FirstDownsPenalty"`
	ThirdDowns        Stat `json:"ThirdDowns"` // 3rd down conversions?
	ThirdDownsAtt     Stat `json:"ThirdDownsAtt"`
	ThirdDownsPct     Stat `json:"ThirdDownsPct"`
	FourthDowns       Stat `json:"FourthDowns"`
	FourthDownsAtt    Stat `json:"FourthDownsAtt"`
	FourthDownsPct    Stat `json:"FourthDownsPct"`
	Penalties         Stat `json:"Penalties"`
	PenaltyYds        Stat `json:"PenaltyYds"` // penalty yards
	OffensePlays      Stat `json:"OffensePlays"`
	OffenseYds        Stat `json:"OffenseYds"`
	OffenseAvgYds     Stat `json:"OffenseAvgYds"`
	TotalTD           Stat `json:"TotalTD"`
}

type PassingStats struct {
	PassAttempts    Stat `json:"PassAttempts"`
	PassCompletions Stat `json:"PassCompletions"`
	PassPct         Stat `json:"PassPct"`
	PassYards       Stat `json:"PassYards"`      // yards or gross
	PassGrossYards  Stat `json:"PassGrossYards"` // ???
	PassNetYards    Stat `json:"PassNetYards"`
	PassAvg         Stat `json:"PassAvg"`
	PassYPA         Stat `json:"PassYardsPerAtt"`
	PassTD          Stat `json:"PassTD"`
	PassTDPct       Stat `json:"PassTDPct"`
	PassInt         Stat `json:"PassInt"`
	PassIntPct      Stat `json:"PassIntPct"`
	PassLng         Stat `json:"PassLng"`
	Pass20Plus      Stat `json:"Pass20Plus"`
	Pass40Plus      Stat `json:"Pass40Plus"`
	PassSacks       Stat `json:"PassSacks"`
	PassSackY       Stat `json:"PassSackY"`
}

func (p PassingStats) Yards() string {
	if len(p.PassYards.Text) != 0 {
		return p.PassYards.Text
	} else {
		return p.PassGrossYards.Text
	}
}

type RushingStats struct {
	RushAttempts Stat `json:"RushAttempts"`
	RushYards    Stat `json:"RushYards"`
	RushAverage  Stat `json:"RushAverage"`
	RushTD       Stat `json:"RushTD"`
	RushLng      Stat `json:"RushLng"`
	Rush20Plus   Stat `json:"Rush20Plus"`
	Rush40Plus   Stat `json:"Rush40Plus"`
	RushFumbles  Stat `json:"RushFumbles"`
}

type ReceivingStats struct {
	Targets    Stat `json:"Targets"`
	Receptions Stat `json:"Receptions"`
	RecYards   Stat `json:"RecYards"`
	RecAvg     Stat `json:"RecAverage"`
	RecTD      Stat `json:"RecTD"`
	RecLng     Stat `json:"RecLng"`
	Rec20Plus  Stat `json:"Rec20Plus"`
	Rec40Plus  Stat `json:"Rec40Plus"`
	RecFumbles Stat `json:"RecFumbles"`
}

type DefenseStats struct {
	TackleSolo     Stat `json:"TackleSolo"`
	TackleTotal    Stat `json:"TackleTotal"`
	TackleAst      Stat `json:"TackleAst"`
	Sacks          Stat `json:"Sacks"`
	SackYds        Stat `json:"SackYds"`
	Safeties       Stat `json:"Safeties"`
	TacklesForLoss Stat `json:"TacklesForLoss"`
	Interceptions  Stat `json:"Interceptions"`
	IntTD          Stat `json:"IntTD"`
	IntYds         Stat `json:"IntYds"`
	IntAverage     Stat `json:"IntAverage"`
	IntLng         Stat `json:"IntLng"`
	PassesDefended Stat `json:"PassesDefended"`
	Stuffs         Stat `json:"Stuffs"`
	StuffYds       Stat `json:"StuffYds"`
}

type KickoffReturnStats struct {
	KRRec    Stat `json:"KrRet"`
	KRYds    Stat `json:"KrYds"`
	KRAvg    Stat `json:"KrAvg"`
	KRLong   Stat `json:"KrLng"`
	KRTD     Stat `json:"KrTD"`
	KR20Plus Stat `json:"Kr20Plus"`
	KR40Plus Stat `json:"Kr40Plus"`
	KRFC     Stat `json:"KrFC"`
	KRFum    Stat `json:"KrFum"`
}

type PuntReturnStats struct {
	PRRet    Stat `json:"PrRet"`
	PRyds    Stat `json:"PrYds"`
	PRAvg    Stat `json:"PrAvg"`
	PRLng    Stat `json:"PrLng"`
	PRTD     Stat `json:"PrTD"`
	PR20Plus Stat `json:"Pr20Plus"`
	PR40Plus Stat `json:"Pr40Plus"`
	PRFC     Stat `json:"PrFC"`
	PRFum    Stat `json:"PrFum"`
}

type KickingStats struct {
	FGBlk        Stat `json:"FgBlk"`
	FGMade       Stat `json:"FgMade"`
	FGAtt        Stat `json:"FgAtt"`
	FGPct        Stat `json:"FgPct"`
	FGMade1_19   Stat `json:"FgMade1_19"`
	FGAtt1_19    Stat `json:"FgAtt1_19"`
	FG1_19Pct    Stat `json:"Fg1_19Pct"`
	FGMade20_29  Stat `json:"FgMade20_29"`
	FGAtt20_29   Stat `json:"FgAtt20_29"`
	FG20_29Pct   Stat `json:"Fg20_29Pct"`
	FGMade30_39  Stat `json:"FgMade30_39"`
	FGAtt30_39   Stat `json:"FgAtt30_39"`
	FG30_39Pct   Stat `json:"Fg30_39Pct"`
	FGMade40_49  Stat `json:"FgMade40_49"`
	FGAtt40_49   Stat `json:"FgAtt40_49"`
	FG40_49Pct   Stat `json:"Fg40_49Pct"`
	FgMade50Plus Stat `json:"FgMade50Plus"`
	FgAtt50Plus  Stat `json:"FgAtt50Plus"`
	Fg50PlusPct  Stat `json:"Fg50PlusPct"`
	FgLng        Stat `json:"FgLng"`
	XpBlk        Stat `json:"XpBlk"`
	XpMade       Stat `json:"XpMade"`
	XpAtt        Stat `json:"XpAtt"`
	XpPct        Stat `json:"XpPct"`
	FgAndXpPts   Stat `json:"FgAndXpPts"`
}

type KickoffStats struct {
	Kickoffs    Stat `json:"Kickoffs"`
	KoYds       Stat `json:"KoYds"`
	KoOOB       Stat `json:"KoOOB"`
	KoAvg       Stat `json:"KoAvg"`
	KoTB        Stat `json:"KoTB"` // touchbacks
	KoPct       Stat `json:"KoPct"`
	KoRet       Stat `json:"KoRet"`
	KoRetYds    Stat `json:"KoRetYds"`
	KoRetAvgYds Stat `json:"KoRetAvgYds"`
	KoTD        Stat `json:"KoTD"`
	KoOS        Stat `json:"KoOS"`
	KoOSR       Stat `json:"KoOSR"`
}

type PuntStats struct {
	Punts       Stat `json:"Punts"`
	PuntYds     Stat `json:"PuntYds"`
	PuntNetYds  Stat `json:"PuntNetYds"`
	PuntLng     Stat `json:"PuntLng"`
	PuntAvg     Stat `json:"PuntAvg"`
	PuntNetAvg  Stat `json:"PuntNetAvg"`
	PuntBlk     Stat `json:"PuntBlk"`
	PuntOOB     Stat `json:"PuntOOB"`
	PuntDown    Stat `json:"PuntDown"`
	PuntIn20    Stat `json:"PuntIn20"`
	PuntIn20Pct Stat `json:"PuntIn20Pct"`
	PuntTB      Stat `json:"PuntTB"` // punting touchbacks
	PuntTBPct   Stat `json:"PuntTBPct"`
	PuntFC      Stat `json:"PuntFC"`
	PuntRet     Stat `json:"PuntRet"`
	PuntRetYds  Stat `json:"PuntRetYds"`
	PuntRetAvg  Stat `json:"PuntRetAvg"`
}

type FumbleStats struct {
	Fumbles     Stat `json:"Fumbles"`
	FumLost     Stat `json:"FumLost"`
	FumForced   Stat `json:"FumForced"`
	FumOwnRec   Stat `json:"FumOwnRec"`
	FumOppRec   Stat `json:"FumOppRec"`
	FumRecYds   Stat `json:"FumRecYds"`
	FumTotalRec Stat `json:"FumTotalRec"`
	FumTD       Stat `json:"FumTD"`
}

type CommonStatList struct {
}

type Stat struct {
	Category     string `json:"@category"`
	Abbreviation string `json:"@abbreviation"`
	Text         string `json:"#text"`
}

type PlayerEntry struct {
	Player Player      `json:"player"`
	Stats  PlayerStats `json:"stats"`
}

func (p PlayerEntry) String() string {
	fn := p.Player.FirstName + " " + p.Player.LastName
	return fmt.Sprintf("%-30s %s, %s", fn, p.Player.Position, p.Stats.String())
}

type Player struct {
	ID           string `json:"ID"`
	LastName     string `json:"LastName"`
	FirstName    string `json:"FirstName"`
	JerseyNumber string `json:"JerseyNumber"`
	Position     string `json:"Position"`
}

type PlayerEntryWrapper struct {
	PlayerEntries []PlayerEntry `json:"playerEntry"`
}

type CumulativePlayerStatsWrapper struct {
	CumulativePlayerStats CumulativePlayerStats `json:"cumulativeplayerstats"`
}

type CumulativePlayerStats struct {
	LastUpdatedOn      string             `json:"lastUpdatedOn"`
	PlayerStatsEntries []PlayerStatsEntry `json:"playerstatsentry"`
}

type PlayerStatsEntry struct {
	PlayerEntry
	Team Team
}
