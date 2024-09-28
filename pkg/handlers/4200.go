package handlers

import (
	"bytes"
	"encoding/binary"
	"github.com/panjf2000/gnet"
	"log"
	"sixserver/pkg/helpers"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x4200(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action) {
	gameVersion := unpackGameVersion(pkt.Data)
	log.Printf("Game version: %d", gameVersion)
	lobbies := getLobbies()
	lenLobbies := uint16(len(lobbies))

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, lenLobbies)
	if err != nil {
		log.Fatalf("Error packing length of lobbies: %v", err)
	}

	// Convertir cada lobby a bytes y concatenarlos
	for _, lobby := range lobbies {
		lobbyBytes, errConv := lobby.ToBytes()
		if errConv != nil {
			log.Fatalf("Error converting lobby to bytes: %v", errConv)
		}
		buf.Write(lobbyBytes)
	}

	data := buf.Bytes()
	err = pes6.SendPacketWithData(conn, 0x4201, data)
	helpers.HandleError(err)

	return
}

func unpackGameVersion(data []byte) uint8 {
	var gameVersion uint8
	err := binary.Read(bytes.NewReader(data[:1]), binary.BigEndian, &gameVersion)
	if err != nil {
		log.Fatalf("Error unpacking game version: %v", err)
	}

	return gameVersion
}

func getLobbies() []types.Lobby {
	return []types.Lobby{
		{
			Name:            "Baltika",
			MaxPlayers:      10,
			Players:         make(map[string]interface{}),
			Rooms:           make(map[string]interface{}),
			TypeStr:         "open",
			TypeCode:        95,
			ShowMatches:     true,
			CheckRosterHash: true,
			RoomOrdinal:     0,
			ChatHistory:     make([]interface{}, 0),
		},
	}
}
