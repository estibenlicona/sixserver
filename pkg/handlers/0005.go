package handlers

import (
	"github.com/panjf2000/gnet"
	"sixserver/pkg/protocols/packet"
)

func Handle0x0005(frame []byte, conn gnet.Conn) {
	pktToSend := packet.ApplyXORKey(frame, 0)
	err := conn.AsyncWrite(pktToSend)
	HandleError(err)
}
