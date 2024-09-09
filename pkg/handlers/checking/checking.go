package checking

import "sixservergo/pkg/dispatcher"

// Crear un dispatcher para el servicio "Checking"
var dispat = dispatcher.NewDispatcher()

// Función que retorna el dispatcher para "Checking"
func GetDispatcher() *dispatcher.Dispatcher {
	dispat.RegisterHandler(0x2008, handle2008)
	dispat.RegisterHandler(0x2006, handle2006)
	dispat.RegisterHandler(0x2005, handle2005)
	dispat.RegisterHandler(0x0003, handle0003)
	dispat.RegisterHandler(0x0005, handle0005)
	dispat.RegisterHandler(0x2200, handle2200)
	return dispat
}
