package handlers

import (
	. "github.com/junglemc/JungleTree/internal/net"
	"github.com/junglemc/JungleTree/internal/net/protocol"
	. "github.com/junglemc/JungleTree/internal/pkg/net/packets"
)

func handshakeSetProtocol(c *Client, p Packet) (err error) {
	pkt := p.(ServerboundHandshakePacket)

	c.Protocol = protocol.Protocol(pkt.NextState)
	c.GameProtocolVersion = pkt.ProtocolVersion
	return
}
