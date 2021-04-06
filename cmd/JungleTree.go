package main

import (
	"github.com/junglemc/net"
	"log"
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
	log.Println("New client! Handshaking...")
	c.Protocol = net.ProtocolHandshake
}

func onClientDisconnect(c *net.Client, err error) {
	log.Println("Client disconnected.")
	if err != nil {
		log.Printf("Client error: %s\n", err)
	}
}

func onClientPacket(c *net.Client, pkt net.Packet) {
	log.Println("Packet recv")
	c.Protocol = net.ProtocolLogin
}
