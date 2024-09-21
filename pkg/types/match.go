package types

import "time"

type Match struct {
	HomeTeamID  uint8
	AwayTeamID  uint8
	HomeProfile Profile
	AwayProfile Profile
	ScoreHome   uint8
	ScoreAway   uint8
	StartTime   time.Time
	HomeExit    uint8
	AwayExit    uint8
	Clock       uint8
	State       MatchState
}
