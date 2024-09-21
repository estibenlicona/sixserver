package tcp

import (
	"github.com/panjf2000/gnet"
	"sixserver/pkg/handlers"
	"sixserver/pkg/types"
)

type EventHandlerFunc func(pkt types.Packet, conn gnet.Conn, config *types.Config) (out []byte, action gnet.Action)

type Dispatcher struct {
	Handlers map[uint16]EventHandlerFunc
	Config   *types.Config
}

func NewDispatcher(config *types.Config) *Dispatcher {
	return &Dispatcher{
		Handlers: map[uint16]EventHandlerFunc{
			0x0003: handlers.Handle0x0003,
			0x2005: handlers.Handle0x2005,
			0x2006: handlers.Handle0x2006,
			0x2008: handlers.Handle0x2008,
			0x2200: handlers.Handle0x2200,
			0x3001: handlers.Handle0x3001,
			0x3003: handlers.Handle0x3003,
			0x3010: handlers.Handle0x3010,
			0x3040: handlers.Handle0x3040,
			0x3050: handlers.Handle0x3050,
			0x3060: handlers.Handle0x3060,
			0x3070: handlers.Handle0x3070,
			0x3080: handlers.Handle0x3080,
			0x3087: handlers.Handle0x3087,
			0x3088: handlers.Handle0x3088,
			0x3089: handlers.Handle0x3089,
			0x308a: handlers.Handle0x308a,
			0x3090: handlers.Handle0x3090,
			0x3100: handlers.Handle0x3100,
			0x3120: handlers.Handle0x3120,
			0x4100: handlers.Handle0x4100,
			0x4102: handlers.Handle0x4102,
			0x4200: handlers.Handle0x4200,
			0x4202: handlers.Handle0x4202,
			0x4210: handlers.Handle0x4210,
			0x4300: handlers.Handle0x4300,
			0x4310: handlers.Handle0x4310,
		},
		Config: config,
	}
}

func (d *Dispatcher) RegisterHandler(event uint16, handler EventHandlerFunc) {
	d.Handlers[event] = handler
}
