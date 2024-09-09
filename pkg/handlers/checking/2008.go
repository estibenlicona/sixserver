package checking

import (
	"net"
	"sixservergo/pkg/protocols/pes6"
	"sixservergo/pkg/protocols/types"
)

func handle2008(conn net.Conn, packet types.Packet) {
	pes6.SendZeros(conn, 0x2009, 4, packet.Header.PacketCount)
	data := append(make([]byte, 4), 0x01, 0x01)
	pes6.SendData(conn, 0x200a, data, packet.Header.PacketCount)
	pes6.SendZeros(conn, 0x200b, 0, packet.Header.PacketCount)
}
