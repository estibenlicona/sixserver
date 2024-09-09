package checking

import (
	"log"
	"net"
	"sixservergo/pkg/protocols/types"
)

func handle0003(conn net.Conn, packet types.Packet) {
	log.Println("User Offline")
}
