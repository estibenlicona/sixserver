package tcp

import (
	"github.com/panjf2000/gnet"
	"log"
	"sixserver/pkg/handlers"
	"sixserver/pkg/protocols/packet"
	"sixserver/pkg/types"
)

type Server struct {
	*gnet.EventServer
	*Dispatcher
	Port    int
	Factory *interface{}
}

func NewServer(port int, config *types.Config) *Server {
	dispatcher := NewDispatcher(config)

	return &Server{
		EventServer: &gnet.EventServer{},
		Dispatcher:  dispatcher,
		Port:        port,
	}
}

func (ls *Server) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("NetworkServer is listening on %s (multi-cores: %t, loops: %d)\n", srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}

func (ls *Server) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Printf("NetworkServer Connection opened: %s\n", c.RemoteAddr().String())
	c.SetContext(&types.ConnectionContext{
		PacketCount: 1,
	})

	return
}

func (ls *Server) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	log.Printf("NetworkServer Connection closed: %s\n", c.RemoteAddr().String())
	return
}

func (ls *Server) React(frame []byte, conn gnet.Conn) (out []byte, action gnet.Action) {
	pkt, err := packet.MakePacket(frame)
	handlers.HandleError(err)

	if pkt.Header.ID == 5 {
		handlers.Handle0x0005(frame, conn)
	} else {
		log.Printf("Packet received: %X\n, PacketCount: %v\n", pkt.Header.ID, pkt.Header.PacketCount)
		if handler, ok := ls.Dispatcher.Handlers[pkt.Header.ID]; ok {
			return handler(pkt, conn, ls.Dispatcher.Config)
		} else {
			return handlers.HandleDefault(pkt, conn)
		}
	}
	return
}
