package handlers

import (
	"github.com/junglemc/net"
	"log"
)

func statusHello(c *net.Client, p net.Packet) {
	// pkt := p.(packet.ServerboundStatusHelloPacket)
	log.Println("Hello!")
}
