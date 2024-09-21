package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Lobby struct {
	Name            string
	MaxPlayers      int
	Players         map[string]interface{}
	Rooms           map[string]interface{}
	TypeStr         string
	TypeCode        int
	ShowMatches     bool
	CheckRosterHash bool
	RoomOrdinal     int
	ChatHistory     []interface{}
}

func (l *Lobby) ToBytes() ([]byte, error) {
	var buf bytes.Buffer

	// Serializa TypeCode
	if err := binary.Write(&buf, binary.BigEndian, uint8(l.TypeCode)); err != nil {
		return nil, fmt.Errorf("error writing TypeCode: %v", err)
	}

	// Serializa Name con padding de ceros hasta 32 bytes
	if _, err := buf.Write(AddPadding(l.Name, 32)); err != nil {
		return nil, fmt.Errorf("error writing Name: %v", err)
	}

	// Serializa la longitud de Players
	if err := binary.Write(&buf, binary.BigEndian, uint16(len(l.Players))); err != nil {
		return nil, fmt.Errorf("error writing Players length: %v", err)
	}

	return buf.Bytes(), nil
}
