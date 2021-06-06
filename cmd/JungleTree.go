package main

import (
	"log"

	"github.com/junglemc/JungleTree/internal/configuration"
	"github.com/junglemc/JungleTree/internal/net"
	"github.com/junglemc/JungleTree/internal/net/handlers"
	"github.com/junglemc/JungleTree/internal/startup"
)

var JungleTreeVersion string

func main() {
	conf := configuration.Config()

	s := net.NewServer(
		conf.Network.IP, conf.Network.Port, conf.JavaEdition.OnlineMode, conf.Network.NetworkCompressionThreshold,
		conf.DebugMode, conf.Verbose, handlers.Handshake, handlers.Status, handlers.Login, handlers.Play, net.Disconnect,
	)

	startup.Init()

	addr := s.Address
	if addr == "" {
		addr = "*"
	}

	log.Printf("Server listening on: %s:%d", addr, s.Port)
	err := s.Listen()
	if err != nil {
		log.Panicln(err)
	}
}
