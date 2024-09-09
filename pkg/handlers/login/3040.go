package login

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"sixservergo/pkg/protocols/pes6"
	"sixservergo/pkg/protocols/types"
)

func handle3040(conn net.Conn, packet types.Packet) {

	//Se deben consultar de la base de datos
	profiles := map[int32]types.Profile{
		12345: {ID: 12345, Name: "Player1", PlayTime: 3600, Points: 600},
		67890: {ID: 67890, Name: "Player2", PlayTime: 5400, Points: 950},
	}

	id := int32(binary.BigEndian.Uint32(packet.Data[0:4]))
	profile, found := profiles[id]
	if !found {
		fmt.Printf("Profile with ID %d not found\n", id)
	}

	var buffer bytes.Buffer
	buffer.Write(make([]byte, 4))

	binary.Write(&buffer, binary.BigEndian, pes6.PadWithZeros(profile.Name, 16))
	padding := make([]byte, 0x18e-20)

	buffer.Write(padding)
	data := buffer.Bytes()

	pes6.SendData(conn, 0x3042, data, packet.Header.PacketCount)
}
