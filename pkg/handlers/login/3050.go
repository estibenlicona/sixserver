package login

import (
	"net"
	"sixservergo/pkg/protocols/pes6"
	"sixservergo/pkg/protocols/types"
)

func handle3050(conn net.Conn, packet types.Packet) {
	pes6.SendZeros(conn, 0x3052, 0x47, packet.Header.PacketCount)
}
