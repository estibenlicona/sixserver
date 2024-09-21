package main

import (
	"fmt"
	"github.com/panjf2000/gnet"
	"log"
	"sixserver/pkg/config"
	"sixserver/pkg/database"
	"sixserver/pkg/tcp"
)

func main() {
	cfg := config.Load()
	database.InitRedis(*cfg)

	loginServer := tcp.NewServer(cfg.LoginPort)
	go func() {
		if err := gnet.Serve(loginServer, fmt.Sprintf("tcp://:%d", loginServer.Port), gnet.WithMulticore(true)); err != nil {
			log.Fatalf("Error starting LoginServer: %v\n", err)
		} else {
			log.Printf("LoginServer is listening on port %d\n", loginServer.Port)
		}
	}()

	lobbyServer := tcp.NewServer(cfg.LobbyPort)
	go func() {
		if err := gnet.Serve(lobbyServer, fmt.Sprintf("tcp://:%d", lobbyServer.Port), gnet.WithMulticore(true)); err != nil {
			log.Fatalf("Error starting LobbyServer: %v\n", err)
		} else {
			log.Printf("LobbyServer is listening on port %d\n", lobbyServer.Port)
		}
	}()

	networkServer := tcp.NewServer(cfg.NetworkPort)
	go func() {
		if err := gnet.Serve(networkServer, fmt.Sprintf("tcp://:%d", networkServer.Port), gnet.WithMulticore(true)); err != nil {
			log.Fatalf("Error starting NetworkServer: %v\n", err)
		} else {
			log.Printf("NetworkServer is listening on port %d\n", networkServer.Port)
		}
	}()

	mainServer := tcp.NewServer(cfg.MainPort)
	go func() {
		if err := gnet.Serve(mainServer, fmt.Sprintf("tcp://:%d", mainServer.Port), gnet.WithMulticore(true)); err != nil {
			log.Fatalf("Error starting CoreServer: %v\n", err)
		} else {
			log.Printf("CoreServer is listening on port %d\n", mainServer.Port)
		}
	}()

	select {}
}
