package pes6

import (
	"github.com/panjf2000/gnet"
	"log"
	"sixserver/pkg/protocols/packet"
	"sixserver/pkg/types"
)

func getConnectionContext(conn gnet.Conn) *types.ConnectionContext {
	return conn.Context().(*types.ConnectionContext)
}

func SendPacketWithData(conn gnet.Conn, id uint16, data []byte) error {
	ctx := getConnectionContext(conn)
	packetReadyToSend := packet.CreatePacketToSend(id, ctx.PacketCount, data)

	err := conn.AsyncWrite(packetReadyToSend)
	if err != nil {
		log.Printf("Error al enviar el paquete: %v", err)
		return err
	}

	ctx.PacketCount++
	return nil
}

func SendPacketWithZeros(conn gnet.Conn, id uint16, size int) error {
	ctx := getConnectionContext(conn)
	data := packet.MakeDataWithOnes(size)
	packetReadyToSend := packet.CreatePacketToSend(id, ctx.PacketCount, data)

	err := conn.AsyncWrite(packetReadyToSend)
	if err != nil {
		log.Printf("Error al enviar el paquete: %v", err)
		return err
	}
	ctx.PacketCount++
	return nil
}
