package handlers

import (
	"bytes"
	"encoding/binary"
	"github.com/panjf2000/gnet"
	"sixserver/pkg/helpers"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x4102(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action) {

	profileId := readProfileId(pkt.Data)
	profile := types.Profile{
		ID: profileId, Name: "Player1", PlayTime: 3600, Points: 600, Rating: 2,
	}

	stats := types.Stats{
		ProfileId: profile.ID,
	}

	profileInfo := packProfileInfo2(profile, stats)
	err := pes6.SendPacketWithData(conn, 0x4103, profileInfo)
	helpers.HandleError(err)

	return
}

func readProfileId(data []byte) uint32 {
	return binary.BigEndian.Uint32(data)
}

func packProfileInfo2(profile types.Profile, stats types.Stats) []byte {

	var buffer bytes.Buffer

	buffer.Write([]byte{0, 0, 0, 0})

	err := binary.Write(&buffer, binary.BigEndian, profile.ID)
	helpers.HandleError(err)

	var name = types.AddPadding(profile.Name, 16)
	buffer.Write(name)

	var division = types.GetDivision(profile.Points)
	err = binary.Write(&buffer, binary.BigEndian, division)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, profile.Points)
	helpers.HandleError(err)

	var matchesPlayed = stats.Wins + stats.Losses + stats.Draws
	err = binary.Write(&buffer, binary.BigEndian, matchesPlayed)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, stats.Wins)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, stats.Losses)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, stats.Draws)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, stats.StreakCurrent)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, stats.StreakBest)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, profile.Disconnects)
	helpers.HandleError(err)

	buffer.Write([]byte{0, 0})

	err = binary.Write(&buffer, binary.BigEndian, stats.GoalsScored)
	helpers.HandleError(err)

	buffer.Write([]byte{0, 0})

	err = binary.Write(&buffer, binary.BigEndian, stats.GoalsAllowed)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, profile.FavTeam)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, profile.FavPlayer)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, profile.Rank)
	helpers.HandleError(err)

	return buffer.Bytes()
}
