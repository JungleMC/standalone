package handlers

import (
	"github.com/junglemc/net"
	"github.com/junglemc/net/protocol"
	"github.com/junglemc/packet"
)

func handshakeSetProtocol(c *net.Client, p net.Packet) (err error) {
	pkt := p.(packet.ServerboundHandshakePacket)

	c.Protocol = protocol.Protocol(pkt.NextState)
	c.GameProtocolVersion = pkt.ProtocolVersion
	return
}
