package types

type UserState struct {
	LobbyID     uint8
	IP1         []byte
	IP2         []byte
	UDPPort1    uint16
	UDPPort2    uint16
	SomeField   uint16
	InRoom      uint8
	NoLobbyChat int32
	Room        interface{}
	TeamID      uint8
}

type User struct {
	Profile Profile
	State   UserState
}

func GetDivision(points uint32) uint8 {
	var maxDivision uint8 = 4
	thresholds := []uint32{250, 450, 600, 750}

	for division, threshold := range thresholds {
		if points < threshold {
			return uint8(division)
		}
	}
	return maxDivision
}
