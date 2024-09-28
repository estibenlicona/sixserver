package handlers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/panjf2000/gnet"
	"sixserver/pkg/helpers"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x4202(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action) {

	var lobbyId uint8
	err := binary.Read(bytes.NewReader(pkt.Data[:4]), binary.BigEndian, &lobbyId)
	helpers.HandleError(err)

	var udpPort1 uint16
	err = binary.Read(bytes.NewReader(pkt.Data[17:19]), binary.BigEndian, &udpPort1)
	helpers.HandleError(err)

	var udpPort2 uint16
	err = binary.Read(bytes.NewReader(pkt.Data[35:37]), binary.BigEndian, &udpPort2)
	helpers.HandleError(err)

	user := types.User{}
	user.State = types.UserState{
		LobbyID:     lobbyId,
		IP1:         pkt.Data[1:17],
		IP2:         pkt.Data[19:35],
		UDPPort1:    udpPort1,
		UDPPort2:    udpPort2,
		InRoom:      0,
		NoLobbyChat: 0,
		Room:        nil,
		TeamID:      0,
	}

	err = pes6.SendPacketWithZeros(conn, 0x4203, 4)
	helpers.HandleError(err)

	var data []byte
	data, err = formatPlayerInfo(user, 0)
	helpers.HandleError(err)

	err = pes6.SendPacketWithData(conn, 0x4220, data)
	helpers.HandleError(err)

	return
}

func formatPlayerInfo(usr types.User, roomId int32) ([]byte, error) {
	var buf bytes.Buffer

	// Empaqueta el ID del perfil del usuario
	if err := binary.Write(&buf, binary.BigEndian, usr.Profile.ID); err != nil {
		return nil, fmt.Errorf("error writing profile ID: %v", err)
	}

	// Empaqueta el nombre del perfil del usuario con padding de ceros hasta 16 bytes
	nameBytes := types.AddPadding(usr.Profile.Name, 16)
	if _, err := buf.Write(nameBytes); err != nil {
		return nil, fmt.Errorf("error writing profile name: %v", err)
	}

	// Empaqueta el estado de InRoom del usuario
	if err := binary.Write(&buf, binary.BigEndian, usr.State.InRoom); err != nil {
		return nil, fmt.Errorf("error writing inRoom state: %v", err)
	}

	// Empaqueta el roomId
	if err := binary.Write(&buf, binary.BigEndian, roomId); err != nil {
		return nil, fmt.Errorf("error writing room ID: %v", err)
	}

	// Empaqueta el estado de NoLobbyChat del usuario
	if err := binary.Write(&buf, binary.BigEndian, usr.State.NoLobbyChat); err != nil {
		return nil, fmt.Errorf("error writing noLobbyChat state: %v", err)
	}

	// Empaqueta dos bytes adicionales con valor 0
	if err := binary.Write(&buf, binary.BigEndian, uint8(0)); err != nil {
		return nil, fmt.Errorf("error writing additional byte 1: %v", err)
	}
	if err := binary.Write(&buf, binary.BigEndian, uint8(0)); err != nil {
		return nil, fmt.Errorf("error writing additional byte 2: %v", err)
	}

	return buf.Bytes(), nil
}
