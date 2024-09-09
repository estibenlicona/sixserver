package types

import "time"

type Match struct {
	homeTeamID  int
	awayTeamID  int
	homeProfile *Profile
	awayProfile *Profile
	scoreHome   int
	scoreAway   int
	startTime   time.Time
	homeExit    *int
	awayExit    *int
}
