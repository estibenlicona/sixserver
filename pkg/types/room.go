package types

type Room struct {
	ID                   uint16
	Name                 string
	MatchTime            uint8
	MatchSettings        interface{}
	UsePassword          bool
	Password             string
	Players              []interface{}
	ReadyCount           uint8
	Owner                interface{}
	Match                Match
	MatchStarter         interface{}
	TeamSelection        interface{}
	Lobby                interface{}
	ParticipatingPlayers []interface{}
	Phase                uint8
}
