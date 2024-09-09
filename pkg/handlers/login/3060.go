package login

import (
	"net"
	"sixservergo/pkg/protocols/pes6"
	"sixservergo/pkg/protocols/types"
)

func handle3060(conn net.Conn, packet types.Packet) {
	pes6.SendZeros(conn, 0x3062, 0, packet.Header.PacketCount)
}
