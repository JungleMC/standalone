package handlers

import (
	"github.com/junglemc/net"
	"github.com/junglemc/net/codec"
	"github.com/junglemc/net/packet"
)

func handshakeSetProtocol(c *net.Client, p codec.Packet) (err error) {
	pkt := p.(packet.ServerboundHandshakePacket)

	c.Protocol = codec.Protocol(pkt.NextState)
	c.GameProtocolVersion = pkt.ProtocolVersion
	return
}
