package login

import (
	"bytes"
	"encoding/binary"
	"net"
	"sixservergo/pkg/protocols/pes6"
	"sixservergo/pkg/protocols/types"
)

func handle308a(conn net.Conn, packet types.Packet) {
	buffer := new(bytes.Buffer)

	value := uint32(0xfffffedd)
	binary.Write(buffer, binary.BigEndian, value)

	data := buffer.Bytes()
	pes6.SendData(conn, 0x3087, data, packet.Header.PacketCount)
}
