package handlers

import (
	"bytes"
	"encoding/binary"
	"github.com/panjf2000/gnet"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x4210(pkt types.Packet, conn gnet.Conn) (out []byte, action gnet.Action) {
	err := pes6.SendPacketWithZeros(conn, 0x4211, 4)
	HandleError(err)

	profile := types.Profile{
		ID: 12345, Name: "Player1", PlayTime: 3600, Points: 600, Rating: 2,
	}

	user := types.User{
		Profile: profile,
	}

	roomId := uint32(0)
	stats := getStats(profile.ID)
	playerInfo := packPlayerInfo(user, roomId, stats)

	err = pes6.SendPacketWithData(conn, 0x4212, playerInfo)
	HandleError(err)

	err = pes6.SendPacketWithZeros(conn, 0x4213, 4)
	HandleError(err)
	return
}

func getStats(profileId uint32) types.Stats {
	stats := types.Stats{
		ProfileId: profileId,
	}
	return stats
}

func packPlayerInfo(user types.User, roomId uint32, stats types.Stats) []byte {
	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, user.Profile.ID)
	HandleError(err)

	buffer.Write(types.AddPadding(user.Profile.Name, 48))

	var groupId uint32 = 0
	err = binary.Write(&buffer, binary.BigEndian, groupId)
	HandleError(err)

	var groupName = "Playmakers"
	buffer.Write(types.AddPadding(groupName, 48))

	var groupMemberStatus uint8 = 1
	err = binary.Write(&buffer, binary.BigEndian, groupMemberStatus)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, types.GetDivision(user.Profile.Points))
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, roomId)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, user.Profile.Points)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, user.Profile.Rating)
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

	buffer.Write([]byte{0, 0, 0})

	return buffer.Bytes()
}

func packProfileInfo(profile types.Profile, stats types.Stats) []byte {

	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, profile.ID)
	HandleError(err)

	buffer.Write(types.AddPadding(profile.Name, 48))

	var groupId uint32 = 0
	err = binary.Write(&buffer, binary.BigEndian, groupId)
	HandleError(err)

	var groupName = "Playmakers"
	buffer.Write(types.AddPadding(groupName, 48))

	var groupMemberStatus uint8 = 1
	err = binary.Write(&buffer, binary.BigEndian, groupMemberStatus)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, types.GetDivision(profile.Points))
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, profile.Points)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, profile.Rating)
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

	err = binary.Write(&buffer, binary.BigEndian, stats.GoalsScored)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, stats.GoalsAllowed)
	HandleError(err)

	var comments = "Fiveserver rules!"
	buffer.Write(types.AddPadding(comments, 256))

	err = binary.Write(&buffer, binary.BigEndian, profile.Rank)
	HandleError(err)

	var competitionGoldMedals = 0
	err = binary.Write(&buffer, binary.BigEndian, competitionGoldMedals)
	HandleError(err)

	var competitionSilverMedals = 0
	err = binary.Write(&buffer, binary.BigEndian, competitionSilverMedals)
	HandleError(err)

	var unknown = 0
	err = binary.Write(&buffer, binary.BigEndian, unknown)
	HandleError(err)

	var cupGoldMedals = 0
	err = binary.Write(&buffer, binary.BigEndian, cupGoldMedals)
	HandleError(err)

	var cupSilverMedals = 0
	err = binary.Write(&buffer, binary.BigEndian, cupSilverMedals)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, unknown)
	HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, unknown)
	HandleError(err)

	var language = 0
	err = binary.Write(&buffer, binary.BigEndian, language)
	HandleError(err)

	var recentTeams = packRecentUsedTeams(stats)
	buffer.Write(recentTeams)

	return buffer.Bytes()
}

func packRecentUsedTeams(stats types.Stats) []byte {
	var buffer bytes.Buffer

	for _, team := range stats.Teams {
		err := binary.Write(&buffer, binary.BigEndian, team)
		HandleError(err)
	}

	return buffer.Bytes()
}
