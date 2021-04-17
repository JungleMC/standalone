package handlers

import (
    "github.com/google/uuid"
    "github.com/junglemc/net"
    "github.com/junglemc/net/auth"
    "github.com/junglemc/net/codec"
    "github.com/junglemc/net/packet"
)

func loginStart(c *net.Client, p codec.Packet) (err error) {
    pkt := p.(packet.ServerboundLoginStartPacket)

    c.Profile.Name = pkt.Username

    if c.Server.OnlineMode {
        pkt := &packet.ClientboundLoginEncryptionRequest{
            ServerId:    "",
            PublicKey:   c.Server.PublicKey(),
            VerifyToken: c.VerifyToken,
        }
        return c.Send(pkt)
    } else {
        err = c.EnableCompression()
        if err != nil {
            return
        }

        c.Profile.ID, _ = uuid.NewRandom()
        return c.LoginSuccess()
    }
}

func loginEncryptionResponse(c *net.Client, p codec.Packet) (err error) {
    pkt := p.(packet.ServerboundLoginEncryptionResponsePacket)

    sharedSecret, err := auth.DecryptLoginResponse(c, pkt.VerifyToken, pkt.SharedSecret)
    if err != nil {
        return
    }

    err = c.EnableEncryption(sharedSecret)
    if err != nil {
        return
    }

    err = c.EnableCompression()
    if err != nil {
        return
    }
    return c.LoginSuccess()
}
