package handlers

import (
	"github.com/junglemc/net"
	"github.com/junglemc/net/protocol"
	packet2 "github.com/junglemc/packet"
)

func handshakeSetProtocol(c *net.Client, p net.Packet) (err error) {
	pkt := p.(packet2.ServerboundHandshakePacket)

	c.Protocol = protocol.Protocol(pkt.NextState)
	c.GameProtocolVersion = pkt.ProtocolVersion
	return
}
