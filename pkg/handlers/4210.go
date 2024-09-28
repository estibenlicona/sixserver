package handlers

import (
	"bytes"
	"encoding/binary"
	"github.com/panjf2000/gnet"
	"sixserver/pkg/helpers"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x4210(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action) {
	err := pes6.SendPacketWithZeros(conn, 0x4211, 4)
	helpers.HandleError(err)

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
	helpers.HandleError(err)

	err = pes6.SendPacketWithZeros(conn, 0x4213, 4)
	helpers.HandleError(err)
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
	helpers.HandleError(err)

	buffer.Write(types.AddPadding(user.Profile.Name, 48))

	var groupId uint32 = 0
	err = binary.Write(&buffer, binary.BigEndian, groupId)
	helpers.HandleError(err)

	var groupName = "Playmakers"
	buffer.Write(types.AddPadding(groupName, 48))

	var groupMemberStatus uint8 = 1
	err = binary.Write(&buffer, binary.BigEndian, groupMemberStatus)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, types.GetDivision(user.Profile.Points))
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, roomId)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, user.Profile.Points)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, user.Profile.Rating)
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

	buffer.Write([]byte{0, 0, 0})

	return buffer.Bytes()
}

func packProfileInfo(profile types.Profile, stats types.Stats) []byte {

	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, profile.ID)
	helpers.HandleError(err)

	buffer.Write(types.AddPadding(profile.Name, 48))

	var groupId uint32 = 0
	err = binary.Write(&buffer, binary.BigEndian, groupId)
	helpers.HandleError(err)

	var groupName = "Playmakers"
	buffer.Write(types.AddPadding(groupName, 48))

	var groupMemberStatus uint8 = 1
	err = binary.Write(&buffer, binary.BigEndian, groupMemberStatus)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, types.GetDivision(profile.Points))
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, profile.Points)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, profile.Rating)
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

	err = binary.Write(&buffer, binary.BigEndian, stats.GoalsScored)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, stats.GoalsAllowed)
	helpers.HandleError(err)

	var comments = "Fiveserver rules!"
	buffer.Write(types.AddPadding(comments, 256))

	err = binary.Write(&buffer, binary.BigEndian, profile.Rank)
	helpers.HandleError(err)

	var competitionGoldMedals = 0
	err = binary.Write(&buffer, binary.BigEndian, competitionGoldMedals)
	helpers.HandleError(err)

	var competitionSilverMedals = 0
	err = binary.Write(&buffer, binary.BigEndian, competitionSilverMedals)
	helpers.HandleError(err)

	var unknown = 0
	err = binary.Write(&buffer, binary.BigEndian, unknown)
	helpers.HandleError(err)

	var cupGoldMedals = 0
	err = binary.Write(&buffer, binary.BigEndian, cupGoldMedals)
	helpers.HandleError(err)

	var cupSilverMedals = 0
	err = binary.Write(&buffer, binary.BigEndian, cupSilverMedals)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, unknown)
	helpers.HandleError(err)

	err = binary.Write(&buffer, binary.BigEndian, unknown)
	helpers.HandleError(err)

	var language = 0
	err = binary.Write(&buffer, binary.BigEndian, language)
	helpers.HandleError(err)

	var recentTeams = packRecentUsedTeams(stats)
	buffer.Write(recentTeams)

	return buffer.Bytes()
}

func packRecentUsedTeams(stats types.Stats) []byte {
	var buffer bytes.Buffer

	for _, team := range stats.Teams {
		err := binary.Write(&buffer, binary.BigEndian, team)
		helpers.HandleError(err)
	}

	return buffer.Bytes()
}
