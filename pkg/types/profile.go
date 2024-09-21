package types

type Profile struct {
	ID          uint32
	Name        string
	PlayTime    uint32
	Points      uint32
	Rating      uint16
	FavTeam     uint16
	FavPlayer   uint32
	Rank        uint32
	Disconnects uint16
}
