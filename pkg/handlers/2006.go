package handlers

import (
	"encoding/binary"
	"github.com/panjf2000/gnet"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
	"time"
)

func Handle0x2006(pkt types.Packet, conn gnet.Conn) (out []byte, action gnet.Action) {
	currentTime := uint32(time.Now().Unix())
	byteSlice := make([]byte, 4)
	binary.BigEndian.PutUint32(byteSlice, currentTime)

	err := pes6.SendPacketWithData(conn, 0x2007, byteSlice)
	HandleError(err)

	return
}
