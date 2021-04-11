package handlers

import (
    "github.com/google/uuid"
    "github.com/junglemc/net"
    "github.com/junglemc/net/packet"
    "log"
)

func loginStart(c *net.Client, p net.Packet) {
    pkt := p.(packet.ServerboundLoginStart)
    c.Username = pkt.Username

    id, _ := uuid.NewRandom()

    response := &packet.ClientboundLoginSuccess{
        Uuid:     id,
        Username: c.Username,
    }
    bin, err := id.MarshalBinary()
    log.Println(bin)
    err = c.Send(response)
    if err != nil {
        log.Println(err)
    }
    c.Protocol = net.ProtocolPlay
}
