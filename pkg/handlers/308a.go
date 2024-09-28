package handlers

import (
	"bytes"
	"encoding/binary"
	"github.com/panjf2000/gnet"
	"sixserver/pkg/helpers"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x308a(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action) {
	buffer := new(bytes.Buffer)

	value := uint32(0xfffffedd)
	err := binary.Write(buffer, binary.BigEndian, value)
	helpers.HandleError(err)

	data := buffer.Bytes()
	err = pes6.SendPacketWithData(conn, 0x3087, data)
	helpers.HandleError(err)

	return
}
