package handlers

import (
	"encoding/binary"
	"github.com/panjf2000/gnet"
	"go.uber.org/zap/buffer"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x4300(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action) {
	err := pes6.SendPacketWithZeros(conn, 0x4301, 4)
	HandleError(err)

	//roomInfo := packRoomInfo()
	//err = pes6.SendPacketWithData(conn, 0x4302, roomInfo)
	//shared.HandleError(err)

	err = pes6.SendPacketWithZeros(conn, 0x4303, 4)
	HandleError(err)

	return
}

func packRoomInfo() []byte {
	room := types.Room{
		ID:    1,
		Name:  "Prueba",
		Phase: 1,
		Match: types.Match{
			State: types.NotStarted,
			Clock: 0,
		},
	}

	var buff buffer.Buffer

	err := binary.Write(&buff, binary.BigEndian, room.ID)
	HandleError(err)

	err = binary.Write(&buff, binary.BigEndian, room.Phase)
	HandleError(err)

	var matchState = uint8(room.Match.State)
	err = binary.Write(&buff, binary.BigEndian, matchState)
	HandleError(err)

	var roomName = types.AddPadding(room.Name, 64)
	err = binary.Write(&buff, binary.BigEndian, roomName)
	HandleError(err)

	err = binary.Write(&buff, binary.BigEndian, room.Match.Clock)
	HandleError(err)

	return buff.Bytes()
}
