package checking

import (
	"fmt"
	"net"
	"sixservergo/pkg/protocols/pes6"
	"sixservergo/pkg/protocols/types"
)

func handle2005(conn net.Conn, packet types.Packet) {
	serverIP := "127.0.0.1"

	servers := []types.Server{
		{TypeID: -1, ServiceID: 2, ServiceName: "LOGIN", ServerIP: serverIP, ServicePort: 20202, NumUsers: 0, SomeValue: 2},
		{TypeID: -1, ServiceID: 3, ServiceName: "Fiveserver", ServerIP: serverIP, ServicePort: 20200, NumUsers: 0, SomeValue: 3},
		{TypeID: -1, ServiceID: 8, ServiceName: "NETWORK_MENU", ServerIP: serverIP, ServicePort: 20201, NumUsers: 0, SomeValue: 8},
	}

	data, err := pes6.PackServers(servers)
	if err != nil {
		fmt.Println(err.Error())
	}

	pes6.SendZeros(conn, 0x2002, 4, packet.Header.PacketCount)
	pes6.SendData(conn, 0x2003, data, packet.Header.PacketCount)
	pes6.SendZeros(conn, 0x2004, 4, packet.Header.PacketCount)

}
