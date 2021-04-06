package handlers

import (
	"github.com/junglemc/net"
	"github.com/junglemc/net/packet"
)

func handshakeHandle(c *net.Client, p net.Packet) {
	pkt := p.(packet.ServerboundHandshakeHelloPacket)

	c.GameProtocolVersion = pkt.ProtocolVersion

	switch pkt.NextState {
	case 1:
		c.Protocol = net.ProtocolStatus
	case 2:
		c.Protocol = net.ProtocolLogin
	}
}
