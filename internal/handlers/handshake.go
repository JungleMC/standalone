package handlers

import (
	"github.com/junglemc/net"
	"github.com/junglemc/net/packet"
	"log"
)

func handshakeHandle(c *net.Client, p net.Packet) {
	pkt := p.(packet.ServerboundHandshakeHelloPacket)

	c.Protocol = net.Protocol(pkt.NextState)
	c.GameProtocolVersion = pkt.ProtocolVersion
	log.Println("Handshake")
}
