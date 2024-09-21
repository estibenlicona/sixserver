package handlers

import (
	"github.com/panjf2000/gnet"
	"log"
	"sixserver/pkg/types"
)

func Handle0x0003(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action) {
	log.Printf("Packet received: %X\n, PacketCount: %v\n", pkt.Header.ID, pkt.Header.PacketCount)
	return
}
