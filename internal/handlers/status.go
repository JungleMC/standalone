package handlers

import (
    "encoding/json"
    "github.com/junglemc/mc"
    "github.com/junglemc/net"
    "github.com/junglemc/net/packet"
    "log"
)

func statusRequest(c *net.Client, p net.Packet) {
    response := mc.ServerListResponse{
        Version: mc.GameVersion{
            Name:     "1.16.5",
            Protocol: 754,
        },
        Players: mc.ServerListPlayers{
            Max:    10,
            Online: 0,
            Sample: []mc.ServerListPlayer{},
        },
        Description: mc.Chat{Text: "A JungleTree Server"},
    }

    data, err := json.Marshal(response)
    if err != nil {
        log.Println(err)
        return
    }

    responsePkt := &packet.ClientboundStatusResponsePacket{Response: string(data)}
    err = c.Send(responsePkt)
    if err != nil {
        log.Println(err)
    }
}

func statusPing(c *net.Client, p net.Packet) {
    err := c.Send(&packet.ClientboundStatusPongPacket{Time: p.(packet.ServerboundStatusPingPacket).Time})
    if err != nil {
        log.Println(err)
    }
    c.Disconnect = true
}
