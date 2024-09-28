package handlers

import (
	"bytes"
	"encoding/binary"
	"github.com/panjf2000/gnet"
	"sixserver/pkg/helpers"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x3010(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action) {

	profiles := []types.Profile{
		{ID: 12345, Name: "Player1", PlayTime: 3600, Points: 600, Rating: 2},
		{ID: 67890, Name: "Player2", PlayTime: 5400, Points: 250, Rating: 2},
	}

	results := []types.Game{
		{Games: 20},
		{Games: 40},
	}

	var buffer bytes.Buffer
	buffer.Write(make([]byte, 4))

	for i, profile := range profiles {

		//Index
		err := binary.Write(&buffer, binary.BigEndian, byte(i))
		if err != nil {
			return nil, 0
		}

		//ID
		err = binary.Write(&buffer, binary.BigEndian, profile.ID)
		if err != nil {
			return nil, 0
		}

		//Name
		name := types.AddPadding(profile.Name, 48)
		err = binary.Write(&buffer, binary.BigEndian, name)
		if err != nil {
			return nil, 0
		}

		//PlayTime
		err = binary.Write(&buffer, binary.BigEndian, profile.PlayTime)
		if err != nil {
			return nil, 0
		}

		//Division
		division := getDivision(profile.Points)
		err = binary.Write(&buffer, binary.BigEndian, division)
		if err != nil {
			return nil, 0
		}

		//Points
		err = binary.Write(&buffer, binary.BigEndian, profile.Points)
		if err != nil {
			return nil, 0
		}

		//Rating

		err = binary.Write(&buffer, binary.BigEndian, profile.Rating)
		if err != nil {
			return nil, 0
		}

		//Games
		err = binary.Write(&buffer, binary.BigEndian, results[i].Games)
		if err != nil {
			return nil, 0
		}

	}

	data := buffer.Bytes()
	err := pes6.SendPacketWithData(conn, 0x3012, data)
	helpers.HandleError(err)

	return
}

func getDivision(points uint32) uint8 {
	var maxDivision uint8 = 4
	thresholds := []uint32{250, 450, 600, 750}

	for division, threshold := range thresholds {
		if points < threshold {
			return uint8(division)
		}
	}
	return maxDivision
}
