package packet

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"sixserver/pkg/types"
)

var XorKey = []byte{0xa6, 0x77, 0x95, 0x7c}

func ApplyXORKey(packet []byte, start int) []byte {
	keySize := len(XorKey)
	b := make([]byte, len(packet))

	for i := range packet {
		b[i] = packet[i] ^ XorKey[(start+i)%keySize]
	}

	return b
}

func MakeHeader(frame []byte) (types.PacketHeader, error) {
	if err := validateMinimumSize(frame, 8); err != nil {
		return types.PacketHeader{}, err
	}

	header := types.PacketHeader{
		ID:          binary.BigEndian.Uint16(frame[0:2]),
		Length:      binary.BigEndian.Uint16(frame[2:4]),
		PacketCount: binary.BigEndian.Uint32(frame[4:8]),
	}
	return header, nil
}

func validateMinimumSize(frame []byte, minSize int) error {
	if len(frame) < minSize {
		return fmt.Errorf("los bytes del encabezado son demasiado cortos, longitud: %d", len(frame))
	}
	return nil
}

func MakePacket(frame []byte) (types.Packet, error) {
	packetXor := ApplyXORKey(frame, 0)

	if err := validateMinimumSize(packetXor, 24); err != nil {
		return types.Packet{}, err
	}

	header, err := MakeHeader(packetXor[0:8])
	if err != nil {
		return types.Packet{}, err
	}

	data, packetMD5, err := extractMD5FromPacket(packetXor, header)
	if err != nil {
		return types.Packet{}, err
	}

	packet := types.Packet{
		Header: header,
		Data:   data,
		MD5:    packetMD5,
	}
	return packet, nil
}

func extractMD5FromPacket(frame []byte, header types.PacketHeader) ([]byte, []byte, error) {
	md5FromPacket := frame[8:24]
	data := frame[24 : 24+int(header.Length)]

	hash := md5.New()
	hash.Write(frame[0:8])
	hash.Write(data)

	md5CalcBytes := hash.Sum(nil)
	if !compareMD5(md5FromPacket, md5CalcBytes) {
		return nil, nil, fmt.Errorf("¡Suma de comprobación MD5 incorrecta! (esperado: %s, obtenido: %s)",
			hex.EncodeToString(md5CalcBytes),
			hex.EncodeToString(md5FromPacket),
		)
	}
	return data, md5FromPacket, nil
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

func headerToBytes(header types.PacketHeader) []byte {
	var buffer bytes.Buffer
	write := func(data interface{}) bool {
		if err := binary.Write(&buffer, binary.BigEndian, data); err != nil {
			return false
		}
		return true
	}

	if !write(header.ID) || !write(header.Length) || !write(header.PacketCount) {
		return nil
	}

	return buffer.Bytes()
}

func packetToBytes(packet types.Packet) []byte {
	var buffer bytes.Buffer

	write := func(data []byte) bool {
		if _, err := buffer.Write(data); err != nil {
			return false
		}
		return true
	}

	headerBytes := headerToBytes(packet.Header)
	if !write(headerBytes) || !write(packet.MD5) || !write(packet.Data) {
		return nil
	}

	return buffer.Bytes()
}

func MakeDataWithOnes(size int) []byte {
	return bytes.Repeat([]byte{0}, size)
}

func CreatePacket(id uint16, packetCount uint32, data []byte) []byte {
	header := types.PacketHeader{
		ID:          id,
		Length:      uint16(len(data)),
		PacketCount: packetCount,
	}

	packet := types.Packet{
		Header: header,
		Data:   data,
		MD5:    makeMD5FromPacket(types.Packet{Header: header, Data: data}),
	}

	return packetToBytes(packet)
}

func CreatePacketToSend(id uint16, packetCount uint32, data []byte) []byte {
	packet := CreatePacket(id, packetCount, data)
	packet = ApplyXORKey(packet, 0)
	return packet
}

func makeMD5FromPacket(packet types.Packet) []byte {
	hash := md5.New()
	write := func(data []byte) {
		if _, err := hash.Write(data); err != nil {
			log.Printf("Error al escribir en el hash: %v", err)
		}
	}

	write(headerToBytes(packet.Header))
	write(packet.Data)

	return hash.Sum(nil)
}

func MakeDataWithServers(servers []types.Server) ([]byte, error) {
	var buf bytes.Buffer

	write := func(data interface{}) error {
		return binary.Write(&buf, binary.BigEndian, data)
	}

	for _, server := range servers {
		if err := write(int32(server.TypeID)); err != nil {
			return nil, fmt.Errorf("error al escribir TypeID: %v", err)
		}
		if err := write(int32(server.ServiceID)); err != nil {
			return nil, fmt.Errorf("error al escribir ServiceID: %v", err)
		}

		if _, err := buf.Write(types.AddPadding(server.ServiceName, 32)); err != nil {
			return nil, fmt.Errorf("error al escribir ServiceName: %v", err)
		}

		if _, err := buf.Write(types.AddPadding(server.ServerIP, 15)); err != nil {
			return nil, fmt.Errorf("error al escribir ServerIP: %v", err)
		}

		if err := write(int16(server.ServicePort)); err != nil {
			return nil, fmt.Errorf("error al escribir ServicePort: %v", err)
		}
		if err := write(int16(server.NumUsers)); err != nil {
			return nil, fmt.Errorf("error al escribir NumUsers: %v", err)
		}
		if err := write(int16(server.SomeValue)); err != nil {
			return nil, fmt.Errorf("error al escribir SomeValue: %v", err)
		}
	}

	return buf.Bytes(), nil
}
