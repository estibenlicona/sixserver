package handlers

import (
	"github.com/panjf2000/gnet"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func HandleDefault(pkt types.Packet, conn gnet.Conn) (out []byte, action gnet.Action) {
	pkt.Header.ID += 1
	err := pes6.SendPacketWithZeros(conn, pkt.Header.ID, 4)
	HandleError(err)
	return
}
