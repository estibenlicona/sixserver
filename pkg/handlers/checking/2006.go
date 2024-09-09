package checking

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"sixservergo/pkg/protocols/pes6"
	"sixservergo/pkg/protocols/types"
	"time"
)

func handle2006(conn net.Conn, packet types.Packet) {
	currentTime := uint32(time.Now().Unix())

	// Empaquetar el tiempo en formato big-endian uint32
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, currentTime)
	if err != nil {
		log.Println("Error al empaquetar el tiempo:", err)
		return
	}

	// Llamar a sendData con el ID 0x2007 y los datos empaquetados
	pes6.SendData(conn, 0x2007, buf.Bytes(), packet.Header.PacketCount)
}
