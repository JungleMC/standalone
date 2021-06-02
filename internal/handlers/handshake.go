package handlers

import (
	"github.com/junglemc/JungleTree/net"
	. "github.com/junglemc/JungleTree/net/protocol"
	. "github.com/junglemc/JungleTree/packet"
)

func handshakeSetProtocol(c *net.Client, p net.Packet) (err error) {
	pkt := p.(ServerboundHandshakePacket)

	c.Protocol = Protocol(pkt.NextState)
	c.GameProtocolVersion = pkt.ProtocolVersion
	return
}
