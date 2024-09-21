package handlers

import (
	"bytes"
	"encoding/binary"
	"github.com/panjf2000/gnet"
	"log"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x4100(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action) {

	profileIndex := unpackProfileIndex(pkt.Data)
	profiles := []types.Profile{
		{ID: 12345, Name: "Player1", PlayTime: 3600, Points: 600},
		{ID: 67890, Name: "Player2", PlayTime: 5400, Points: 600},
	}
	profile := profiles[profileIndex]

	data := createData(profile.ID)
	data = appendData(data)

	err := pes6.SendPacketWithData(conn, 0x4101, data)
	HandleError(err)

	return
}

func unpackProfileIndex(data []byte) uint8 {
	var profileIndex uint8
	err := binary.Read(bytes.NewReader(data[:1]), binary.BigEndian, &profileIndex)
	if err != nil {
		log.Fatalf("Error unpacking profile index: %v", err)
	}
	return profileIndex
}

func createData(profileID uint32) []byte {
	// Crea un buffer de bytes
	buf := new(bytes.Buffer)

	// Escribe los primeros 4 bytes como ceros
	_, err := buf.Write([]byte{0, 0, 0, 0})
	if err != nil {
		log.Fatalf("Error writing zeros: %v", err)
	}

	// Empaqueta el profileID en big-endian y escribe en el buffer
	err = binary.Write(buf, binary.BigEndian, profileID)
	if err != nil {
		log.Fatalf("Error packing profile ID: %v", err)
	}

	return buf.Bytes()
}

func appendData(data []byte) []byte {
	// Crea un buffer de bytes
	buf := bytes.NewBuffer(data)

	// Añade los bytes especificados
	_, err := buf.Write([]byte{
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 7 bytes de 0xff
		0x80,                                                                                     // 1 byte de 0x80
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 15 bytes de 0xff
		0xc0,                                                 // 1 byte de 0xc0
		0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x01, 0x00, // 9 bytes específicos
	})
	if err != nil {
		log.Fatalf("Error appending data: %v", err)
	}

	return buf.Bytes()
}
