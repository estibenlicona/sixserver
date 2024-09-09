package main

import (
	"log"
	"sixservergo/pkg/config"
	"sixservergo/pkg/dispatcher"
	"sixservergo/pkg/handlers/checking"
	"sixservergo/pkg/handlers/login"
	"sixservergo/pkg/tcp"
)

func main() {
	cfg := config.Load()
	go startTCPService("Checking Service", cfg.CheckingTCPPort, checking.GetDispatcher())
	go startTCPService("Login Service", cfg.LoginTCPPort, login.GetDispatcher())
	select {}
}

func startTCPService(serviceName string, port int, dispatcher *dispatcher.Dispatcher) {
	if err := tcp.Start(port, dispatcher); err != nil {
		log.Fatalf("Error starting %s: %v", serviceName, err)
	}
}
