package login

import (
	"net"
	"sixservergo/pkg/protocols/pes6"
	"sixservergo/pkg/protocols/types"
)

func handle3070(conn net.Conn, packet types.Packet) {
	pes6.SendZeros(conn, 0x3071, 4, packet.Header.PacketCount)
	pes6.SendZeros(conn, 0x3073, 4, packet.Header.PacketCount)
}
