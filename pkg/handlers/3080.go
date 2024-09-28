package handlers

import (
	"github.com/panjf2000/gnet"
	"sixserver/pkg/helpers"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x3080(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action) {
	err := pes6.SendPacketWithZeros(conn, 0x3082, 4)
	helpers.HandleError(err)

	err = pes6.SendPacketWithZeros(conn, 0x3086, 0)
	helpers.HandleError(err)
	return
}
