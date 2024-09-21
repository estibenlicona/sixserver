package handlers

import (
	"github.com/panjf2000/gnet"
	"sixserver/pkg/types"
	"time"
)

func Handle0x3088(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action) {
	if pkt.Data[2] == 3 {

	} else {

	}

	time.Sleep(2 * time.Second)

	return
}
