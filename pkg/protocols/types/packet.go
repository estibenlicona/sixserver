package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

type Packet struct {
	Header PacketHeader
	Data   []byte
	MD5    []byte
}

func (packet Packet) String() string {
	return fmt.Sprintf("Packet(header: %v, md5: %s, data: %x)",
		packet.Header,
		hex.EncodeToString(packet.MD5),
		packet.Data)
}

func (packet Packet) ToBytes() []byte {
	var buffer bytes.Buffer
	headerBytes := packet.Header.ToBytes()
	buffer.Write(headerBytes)
	buffer.Write(packet.MD5)
	buffer.Write(packet.Data)
	return buffer.Bytes()
}
