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

func (header PacketHeader) ToBytes() []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, header.ID)
	if err != nil {
		fmt.Printf("Error writing ID: %v\n", err)
	}

	err = binary.Write(buffer, binary.BigEndian, header.Length)
	if err != nil {
		fmt.Printf("Error writing ID: %v\n", err)
	}

	err = binary.Write(buffer, binary.BigEndian, header.PacketCount)
	if err != nil {
		fmt.Printf("Error writing ID: %v\n", err)
	}

	return buffer.Bytes()
}
