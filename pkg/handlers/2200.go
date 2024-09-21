package handlers

import (
	"github.com/panjf2000/gnet"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x2200(pkt types.Packet, conn gnet.Conn) (out []byte, action gnet.Action) {
	err := pes6.SendPacketWithZeros(conn, 0x2201, 4)
	HandleError(err)

	err = pes6.SendPacketWithZeros(conn, 0x2203, 4)
	HandleError(err)

	return
}
