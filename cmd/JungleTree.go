package main

import (
	"github.com/junglemc/JungleTree/internal/handlers"
	"github.com/junglemc/net"
	"log"
	"reflect"
)

func main() {
	// TODO: Cleanup manual testing shite
	srv := &net.Server{
		Address:            "0.0.0.0",
		Port:               25565,
		OnClientConnect:    onClientConnect,
		OnClientDisconnect: onClientDisconnect,
		OnClientPacket:     onClientPacket,
	}

	run(srv)
}

func run(srv *net.Server) {
	err := srv.Listen()
	if err != nil {
		log.Fatalln(err)
	}
}

func onClientConnect(c *net.Client) {
	log.Printf("Client connected: %s\n", c.Connection.RemoteAddr().String())
	c.Protocol = net.ProtocolHandshake
}

func onClientDisconnect(c *net.Client, err error) {
	log.Printf("Client disconnect: %s\n", c.Connection.RemoteAddr().String())
	if err != nil {
		if err.Error() == "EOF" {
			return
		}
		log.Printf("Client error: %s\n", err)
	}
}

func onClientPacket(c *net.Client, pkt net.Packet) {
	var find map[reflect.Type]func(c *net.Client, pkt net.Packet)

	switch c.Protocol {
	case net.ProtocolHandshake:
		find = handlers.HandshakeHandlers
	}

	if find == nil {
		panic("not implemented")
	}

	funcCall := find[reflect.TypeOf(pkt)]
	if pkt == nil {
		panic("not found") // TODO: Cleanup
	}
	funcCall(c, pkt)
	log.Println(c.GameProtocolVersion)
}
