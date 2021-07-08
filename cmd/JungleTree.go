package main

import (
	"github.com/junglemc/JungleTree/internal/services"
	"log"

	"github.com/junglemc/JungleTree/internal/configuration"
	"github.com/junglemc/JungleTree/internal/net"
	"github.com/junglemc/JungleTree/internal/net/handlers"
	"github.com/junglemc/JungleTree/internal/startup"
	"github.com/junglemc/JungleTree/pkg/event"
)

func main() {
	go func() {
		err := services.World("", 50051)
		if err != nil {
			panic(err)
		}
	}()

	conf := configuration.Config()

	s := net.NewServer(
		conf.Network.IP, conf.Network.Port, conf.JavaEdition.OnlineMode, conf.Network.NetworkCompressionThreshold,
		conf.DebugMode, conf.Verbose, handlers.Handshake, handlers.Status, handlers.Login, handlers.Play, net.Disconnect,
	)

	startup.Init()
	defer rpcClose()

	addr := s.Address
	if addr == "" {
		addr = "*"
	}

	log.Printf("Server listening on: %s:%d", addr, s.Port)
	err := s.Listen()
	if err != nil {
		log.Panicln(err)
	}

	event.Trigger(event.ServerLoadedEvent{})
}

func rpcClose() {
	startup.WorldConnection.Close()
}
