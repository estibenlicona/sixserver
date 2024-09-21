package handlers

import (
	"github.com/panjf2000/gnet"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x2008(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action) {
	var err = pes6.SendPacketWithZeros(conn, 0x2009, 4)
	HandleError(err)

	err = pes6.SendPacketWithData(conn, 0x200a, []byte{0x01, 0x01})
	HandleError(err)

	err = pes6.SendPacketWithZeros(conn, 0x200b, 4)
	HandleError(err)

	return
}
