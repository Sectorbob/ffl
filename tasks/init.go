package tasks

import (
	"time"
)

var NewYork *time.Location

func init() {
	var err error
	NewYork, err = time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
}
