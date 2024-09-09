package login

import "sixservergo/pkg/dispatcher"

// Crear un dispatcher para el servicio "Checking"
var dispat = dispatcher.NewDispatcher()

// Función que retorna el dispatcher para "Checking"
func GetDispatcher() *dispatcher.Dispatcher {
	dispat.RegisterHandler(0x3001, handle3001)
	dispat.RegisterHandler(0x3003, handle3003)
	dispat.RegisterHandler(0x3010, handle3010)
	dispat.RegisterHandler(0x3060, handle3060)
	dispat.RegisterHandler(0x3040, handle3040)
	dispat.RegisterHandler(0x3050, handle3050)
	dispat.RegisterHandler(0x3070, handle3070)
	dispat.RegisterHandler(0x308a, handle308a)
	return dispat
}
