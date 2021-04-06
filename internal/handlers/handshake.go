package handlers

import (
    "github.com/junglemc/mc/packet"
    "github.com/junglemc/net"
)

func handshakeSetProtocol(c *net.Client, p net.Packet) {
    pkt := p.(packet.ServerboundHandshakeSetProtocol)

    c.Protocol = net.Protocol(pkt.NextState)
    c.GameProtocolVersion = pkt.ProtocolVersion
}
