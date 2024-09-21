package handlers

import (
	"bytes"
	"encoding/binary"
	"github.com/panjf2000/gnet"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x4102(pkt types.Packet, conn gnet.Conn) (out []byte, action gnet.Action) {

	profileId := readProfileId(pkt.Data)
	profile := types.Profile{
		ID: profileId, Name: "Player1", PlayTime: 3600, Points: 600, Rating: 2,
	}

	stats := types.Stats{
		ProfileId: profile.ID,
	}

	profileInfo := packProfileInfo2(profile, stats)
	err := pes6.SendPacketWithData(conn, 0x4103, profileInfo)
	HandleError(err)

	return
}

func readProfileId(data []byte) uint32 {
	return binary.BigEndian.Uint32(data)
}

func packProfileInfo2(profile types.Profile, stats types.Stats) []byte {

	var buffer bytes.Buffer

	buffer.Write([]byte{0, 0, 0, 0})

	err := binary.Write(&buffer, binary.BigEndian, profile.ID)
	HandleError(err)

	var name = types.AddPadding(profile.Name, 16)
	buffer.Write(name)

	var division = types.GetDivision(profile.Points)
	err = binary.Write(&buffer, binary.BigEndian, division)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, profile.Points)
	HandleError(err)

	var matchesPlayed = stats.Wins + stats.Losses + stats.Draws
	err = binary.Write(&buffer, binary.BigEndian, matchesPlayed)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, stats.Wins)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, stats.Losses)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, stats.Draws)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, stats.StreakCurrent)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, stats.StreakBest)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, profile.Disconnects)
	HandleError(err)

	buffer.Write([]byte{0, 0})

	err = binary.Write(&buffer, binary.BigEndian, stats.GoalsScored)
	HandleError(err)

	buffer.Write([]byte{0, 0})

	err = binary.Write(&buffer, binary.BigEndian, stats.GoalsAllowed)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, profile.FavTeam)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, profile.FavPlayer)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, profile.Rank)
	HandleError(err)

	return buffer.Bytes()
}
