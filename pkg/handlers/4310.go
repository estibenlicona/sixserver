package handlers

import (
	"github.com/panjf2000/gnet"
	"log"
	"sixserver/pkg/types"
)

func Handle0x4310(pkt types.Packet, conn gnet.Conn) (out []byte, action gnet.Action) {
	roomName := types.RemovePadding(pkt.Data[0:64])
	lobby := types.Lobby{}

	room := types.Room{
		ID:          1,
		Name:        string(roomName),
		Lobby:       lobby,
		UsePassword: false,
	}

	log.Printf("Room name: %s", room.Name)

	return
}
