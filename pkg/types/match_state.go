package types

type MatchState int

const (
	NotStarted MatchState = iota
	FirstHalf
	HalfTime
	SecondHalf
	BeforeExtraTime
	EtFirstHalf
	EtBreak
	EtSecondHalf
	BeforePenalties
	Penalties
	Finished
)

var StateText = map[MatchState]string{
	NotStarted:      "Not started",
	FirstHalf:       "1st half",
	HalfTime:        "Half-time",
	SecondHalf:      "2nd half",
	BeforeExtraTime: "Normal time finished",
	EtFirstHalf:     "Extra-time 1st half",
	EtBreak:         "Extra-time intermission",
	EtSecondHalf:    "Extra-time 2nd half",
	BeforePenalties: "Before penalties",
	Penalties:       "Penalties",
	Finished:        "Finished",
}
