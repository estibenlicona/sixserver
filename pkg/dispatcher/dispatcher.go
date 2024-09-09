package dispatcher

import (
	"net"
	"sixservergo/pkg/protocols/types"
)

// Definimos un tipo de handler como una función
type HandlerFunc func(net.Conn, types.Packet)

// Dispatcher maneja un conjunto de handlers para un servicio específico
type Dispatcher struct {
	handlers map[uint16]HandlerFunc
}

// NewDispatcher crea un nuevo dispatcher vacío
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		handlers: make(map[uint16]HandlerFunc),
	}
}

// RegisterHandler registra un handler para un ID específico
func (d *Dispatcher) RegisterHandler(id uint16, handler HandlerFunc) {
	d.handlers[id] = handler
}

// DispatchPacket despacha un paquete al handler correcto
func (d *Dispatcher) DispatchPacket(conn net.Conn, packet types.Packet) {
	handler, exists := d.handlers[packet.Header.ID]
	if exists {
		handler(conn, packet)
	}
}
