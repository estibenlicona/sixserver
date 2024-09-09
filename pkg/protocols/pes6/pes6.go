package pes6

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"sixservergo/pkg/protocols/types"
	"sync/atomic"
)

func PackServers(servers []types.Server) ([]byte, error) {
	var buf bytes.Buffer

	for _, server := range servers {
		if err := binary.Write(&buf, binary.BigEndian, int32(server.TypeID)); err != nil {
			return nil, fmt.Errorf("error al escribir TypeID: %v", err)
		}
		if err := binary.Write(&buf, binary.BigEndian, int32(server.ServiceID)); err != nil {
			return nil, fmt.Errorf("error al escribir ServiceID: %v", err)
		}

		serviceNameBytes := PadWithZeros(server.ServiceName, 32)
		if _, err := buf.Write(serviceNameBytes); err != nil {
			return nil, fmt.Errorf("error al escribir ServiceName: %v", err)
		}

		ipBytes := PadWithZeros(server.ServerIP, 15)
		if _, err := buf.Write(ipBytes); err != nil {
			return nil, fmt.Errorf("error al escribir ServerIP: %v", err)
		}

		if err := binary.Write(&buf, binary.BigEndian, int16(server.ServicePort)); err != nil {
			return nil, fmt.Errorf("error al escribir ServicePort: %v", err)
		}

		if err := binary.Write(&buf, binary.BigEndian, int16(server.NumUsers)); err != nil {
			return nil, fmt.Errorf("error al escribir NumUsers: %v", err)
		}
		if err := binary.Write(&buf, binary.BigEndian, int16(server.SomeValue)); err != nil {
			return nil, fmt.Errorf("error al escribir SomeValue: %v", err)
		}
	}

	return buf.Bytes(), nil
}

func ProcessPacketHeader(buffer []byte, headerSize int, packetCount *uint32) (types.PacketHeader, error) {
	bufferHeader := buffer[:headerSize]
	decryptedHeader := xorData(bufferHeader, 0)
	packetHeader, err := MakePacketHeader(decryptedHeader, packetCount) //Revisar si funcion MakePacketHeader
	if err != nil {
		return types.PacketHeader{}, err
	}
	return packetHeader, nil
}

func CalculateTotalPacketSize(header types.PacketHeader) int {
	return int(header.Length) + 24
}

func ProcessCompletePacket(buffer []byte, totalPacketSize int, packetCount *uint32) (types.Packet, error) {
	packetData := buffer[:totalPacketSize]
	decryptedPacketData := xorData(packetData, 0)
	packet, err := MakePacket(decryptedPacketData, packetCount) // Revisar si se puede cambiar MakePacket por CreateNewPacket
	if err != nil {
		return types.Packet{}, err
	}
	return packet, nil
}

func SendZeros(conn net.Conn, id uint16, length int, packetCount *uint32) {
	data := make([]byte, length)
	SendData(conn, id, data, packetCount)
}

func SendData(conn net.Conn, id uint16, data []byte, packetCount *uint32) {
	header := types.PacketHeader{
		ID:          id,
		Length:      uint16(len(data)),
		PacketCount: packetCount,
	}

	packet := CreateNewPacket(header, data)

	Send(conn, packet)
}

func Send(conn net.Conn, packet types.Packet) {

	pktBytes := packet.ToBytes()
	xoredData := xorData(pktBytes, 0)

	n, err := conn.Write(xoredData)
	if err != nil {
		log.Println("Error enviando datos:", err)
	} else {
		log.Printf("Se enviaron %d bytes correctamente", n)
	}

	atomic.AddUint32(packet.Header.PacketCount, 1)
}

var XOR_KEY = []byte{0xa6, 0x77, 0x95, 0x7c}

func xorData(data []byte, start int) []byte {
	keySize := len(XOR_KEY)
	result := make([]byte, len(data))

	for i, c := range data {
		result[i] = c ^ XOR_KEY[(start+i)%keySize]
	}

	return result
}

func HandlerError(err error) bool {
	var hasError bool = err != nil

	if hasError {
		log.Fatal(err)
	}

	return hasError
}

func compareMD5(md5a, md5b []byte) bool {
	if len(md5a) != len(md5b) {
		return false
	}
	for i := range md5a {
		if md5a[i] != md5b[i] {
			return false
		}
	}
	return true
}

func PadWithZeros(value string, total int) []byte {
	bytes := []byte(value)
	if len(bytes) > total {
		return bytes[:total]
	}

	padding := make([]byte, total-len(bytes))
	return append(bytes, padding...)
}

func StripZeros(b []byte) []byte {
	return bytes.TrimRight(b, "\x00")
}

func CreateNewPacket(header types.PacketHeader, data []byte) types.Packet {
	packet := types.Packet{
		Header: header,
		Data:   data,
	}

	lenData := uint16(len(data))
	header.Length = lenData
	headerBytes := header.ToBytes()

	hasher := md5.New()
	hasher.Write(headerBytes)
	hasher.Write(data)
	packet.MD5 = hasher.Sum(nil)
	return packet
}

func MakePacketHeader(bs []byte, packetCount *uint32) (types.PacketHeader, error) {
	if len(bs) < 8 {
		return types.PacketHeader{}, fmt.Errorf("el buffer es demasiado pequeño para un encabezado")
	}

	id := binary.BigEndian.Uint16(bs[0:2])

	length := binary.BigEndian.Uint16(bs[2:4])

	return types.PacketHeader{
		ID:          id,
		Length:      length,
		PacketCount: packetCount,
	}, nil
}

func MakePacket(bs []byte, packetCount *uint32) (types.Packet, error) {

	header, err := MakePacketHeader(bs[0:8], packetCount)
	if err != nil {
		return types.Packet{}, fmt.Errorf("error al crear el encabezado: %v", err)
	}

	md5FromPacket := bs[8:24]

	data := bs[24 : 24+int(header.Length)]

	hash := md5.New()
	hash.Write(bs[0:8])
	hash.Write(data)

	md5CalcBytes := hash.Sum(nil)
	if !compareMD5(md5FromPacket, md5CalcBytes) {
		return types.Packet{}, fmt.Errorf("wrong MD5-checksum! (expected: %s, got: %s)",
			hex.EncodeToString(md5CalcBytes),
			hex.EncodeToString(md5FromPacket),
		)
	}

	p := types.Packet{
		Header: header,
		Data:   data,
		MD5:    md5CalcBytes,
	}

	return p, nil
}
