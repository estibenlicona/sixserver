package handlers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/panjf2000/gnet"
	"sixserver/pkg/helpers"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x3040(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action) {

	profiles := map[int32]types.Profile{
		12345: {ID: 12345, Name: "Player1", PlayTime: 3600, Points: 600},
		67890: {ID: 67890, Name: "Player2", PlayTime: 5400, Points: 950},
	}

	id := int32(binary.BigEndian.Uint32(pkt.Data[0:4]))
	profile, found := profiles[id]
	if !found {
		fmt.Printf("Profile with ID %d not found\n", id)
	}

	var buffer bytes.Buffer
	buffer.Write(make([]byte, 4))

	err := binary.Write(&buffer, binary.BigEndian, types.AddPadding(profile.Name, 16))
	helpers.HandleError(err)
	padding := make([]byte, 0x18e-20)

	buffer.Write(padding)
	data := buffer.Bytes()

	err = pes6.SendPacketWithData(conn, 0x3042, data)
	helpers.HandleError(err)

	return
}
