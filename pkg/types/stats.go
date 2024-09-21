package types

type Stats struct {
	ProfileId     uint32
	Wins          uint16
	Losses        uint16
	Draws         uint16
	GoalsScored   uint16
	GoalsAllowed  uint16
	StreakCurrent uint16
	StreakBest    uint16
	Teams         []interface{}
}
