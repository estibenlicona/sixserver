package checking

import (
	"net"
	"sixservergo/pkg/protocols/pes6"
	"sixservergo/pkg/protocols/types"
)

func handle0005(conn net.Conn, packet types.Packet) {
	pes6.Send(conn, packet)
}
