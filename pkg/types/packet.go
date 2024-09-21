package types

import "bytes"

type Packet struct {
	Header PacketHeader
	Data   []byte
	MD5    []byte
}

func AddPadding(value string, total int) []byte {
	bs := []byte(value)
	if len(bs) > total {
		return bs[:total]
	}

	padding := make([]byte, total-len(bs))
	return append(bs, padding...)
}

func RemovePadding(s []byte) []byte {
	isZero := bytes.IndexByte(s, 0)
	if isZero >= 0 {
		return s[:isZero]
	}
	return s
}
