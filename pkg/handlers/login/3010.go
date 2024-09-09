package login

import (
	"bytes"
	"encoding/binary"
	"net"
	"sixservergo/pkg/protocols/pes6"
	"sixservergo/pkg/protocols/types"
)

func handle3010(conn net.Conn, packet types.Packet) {

	profiles := []types.Profile{
		{ID: 12345, Name: "Player1", PlayTime: 3600, Points: 600},
		{ID: 67890, Name: "Player2", PlayTime: 5400, Points: 950},
	}

	results := []types.Game{
		{Games: 10},
		{Games: 20},
	}

	var buffer bytes.Buffer
	buffer.Write(make([]byte, 4))

	for i, profile := range profiles {
		binary.Write(&buffer, binary.BigEndian, byte(i))
		binary.Write(&buffer, binary.BigEndian, profile.ID)
		binary.Write(&buffer, binary.BigEndian, pes6.PadWithZeros(profile.Name, 16))
		binary.Write(&buffer, binary.BigEndian, profile.PlayTime)
		binary.Write(&buffer, binary.BigEndian, getDivision(profile.Points))
		binary.Write(&buffer, binary.BigEndian, profile.Points)
		binary.Write(&buffer, binary.BigEndian, results[i].Games)
	}

	data := buffer.Bytes()
	pes6.SendData(conn, 0x3012, data, packet.Header.PacketCount)
}

func getDivision(points int32) int32 {
	maxDivision := 4
	thresholds := []int32{250, 450, 600, 750}
	for division, threshold := range thresholds {
		if points < threshold {
			return int32(division)
		}
	}
	return int32(maxDivision)
}
