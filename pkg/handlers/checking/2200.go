package checking

import (
	"net"
	"sixservergo/pkg/protocols/pes6"
	"sixservergo/pkg/protocols/types"
)

func handle2200(conn net.Conn, packet types.Packet) {
	pes6.SendZeros(conn, 0x2201, 4, packet.Header.PacketCount)
	pes6.SendZeros(conn, 0x2203, 4, packet.Header.PacketCount)
}
