package handlers

import (
	"github.com/panjf2000/gnet"
	"net"
	"sixserver/pkg/helpers"
	"sixserver/pkg/protocols/packet"
	"sixserver/pkg/protocols/pes6"
	"sixserver/pkg/types"
)

func Handle0x2005(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action) {
	remoteAddr := conn.RemoteAddr().String()
	serverIP, _, err := net.SplitHostPort(remoteAddr)
	helpers.HandleError(err)

	servers := []types.Server{
		{TypeID: -1, ServiceID: 2, ServiceName: "LOGIN", ServerIP: serverIP, ServicePort: 20202, NumUsers: 0, SomeValue: 2},
		{TypeID: -1, ServiceID: 3, ServiceName: "BALTIKA", ServerIP: serverIP, ServicePort: 20200, NumUsers: 0, SomeValue: 3},
		{TypeID: -1, ServiceID: 8, ServiceName: "NETWORK_MENU", ServerIP: serverIP, ServicePort: 20201, NumUsers: 0, SomeValue: 8},
	}

	err = pes6.SendPacketWithZeros(conn, 0x2002, 4)
	helpers.HandleError(err)

	data, _ := packet.MakeDataWithServers(servers)
	err = pes6.SendPacketWithData(conn, 0x2003, data)
	helpers.HandleError(err)

	err = pes6.SendPacketWithZeros(conn, 0x2004, 4)
	helpers.HandleError(err)
	return
}
