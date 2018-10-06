package model

import (
	"time"
)

type Season struct {
	ID   uint32
	Year uint32
}

type Week struct {
	ID       uint32
	Number   int
	Type     string
	SeasonID uint32
	Created  time.Time
	Updated  time.Time
}

type Team struct {
	ID       uint32
	Abbrv    string
	City     string
	Name     string
	MSFID    uint32
	MSFAbbrv string
}

type Game struct {
	ID              uint32
	WeekID          uint32
	HomeID          uint32
	AwayID          uint32
	StartTS         *time.Time
	OriginalStartTS *time.Time
	MSFID           int
	Location        string
	Created         time.Time
	Updated         time.Time
}

type GameResult struct {
	ID        uint32
	GameID    uint32
	HomeScore uint32
	AwayScore uint32
	Created   time.Time
	Updated   time.Time
}

type Player struct {
}
