package handlers

import (
    "github.com/junglemc/net"
    "github.com/junglemc/net/packet"
)

func handshakeSetProtocol(c *net.Client, p net.Packet) {
    pkt := p.(packet.ServerboundHandshakeSetProtocol)

    c.Protocol = net.Protocol(pkt.NextState)
    c.GameProtocolVersion = pkt.ProtocolVersion
}
