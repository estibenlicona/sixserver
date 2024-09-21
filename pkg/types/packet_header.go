package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type PacketHeader struct {
	ID          uint16
	Length      uint16
	PacketCount uint32
}

func (header PacketHeader) String() string {
	return fmt.Sprintf("PacketHeader(ID: %#x, Length: %d, PacketCount: %d)", header.ID, header.Length, header.PacketCount)
}

func (packetHeader PacketHeader) ToBytes() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, packetHeader.ID)
	binary.Write(buffer, binary.BigEndian, packetHeader.Length)
	binary.Write(buffer, binary.BigEndian, packetHeader.PacketCount)
	return buffer.Bytes()
}
