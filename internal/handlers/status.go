package handlers

import (
    "encoding/json"
    "github.com/junglemc/JungleTree/internal/player"
    "github.com/junglemc/JungleTree/pkg"
    "github.com/junglemc/mc/chat"
    "github.com/junglemc/mc/status"
    "github.com/junglemc/net"
    "github.com/junglemc/net/codec"
    "github.com/junglemc/net/packet"
)

func statusRequest(c *net.Client, p codec.Packet) (err error) {
    response := status.ServerListResponse{
        Version: status.GameVersion{
            Name:     "1.16.5",
            Protocol: 754,
        },
        Players: status.ServerListPlayers{
            Max:    pkg.Config().MaxOnlinePlayers,
            Online: player.GetOnlinePlayers(),
            Sample: []status.ServerListPlayer{},
        },
        Description: chat.Message{Text: "A JungleTree Server"},
    }

    data, err := json.Marshal(response)
    if err != nil {
        return
    }

    responsePkt := &packet.ClientboundStatusResponsePacket{Response: string(data)}
    return c.Send(responsePkt)
}

func statusPing(c *net.Client, p codec.Packet) (err error) {
    return c.Send(&packet.ClientboundStatusPongPacket{Time: p.(packet.ServerboundStatusPingPacket).Time})
}
