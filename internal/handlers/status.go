package handlers

import (
    "encoding/json"
    "github.com/google/uuid"
    "github.com/junglemc/mc"
    "github.com/junglemc/mc/packet"
    "github.com/junglemc/net"
    "log"
)

func statusRequest(c *net.Client, p net.Packet) {
    response := ServerListResponse{
        Version: GameVersion{
            Name:     "1.16.5",
            Protocol: 754,
        },
        Players: ServerListPlayers{
            Max:    10,
            Online: 0,
            Sample: []ServerListPlayer{},
        },
        Description: mc.Chat{Text: "A JungleTree Server"},
    }

    data, err := json.Marshal(response)
    if err != nil {
        log.Println(err)
        return
    }

    responsePkt := &packet.ClientboundStatusServerInfo{Response: string(data)}
    err = c.Send(responsePkt)
    if err != nil {
        log.Println(err)
    }
}

func statusPing(c *net.Client, p net.Packet) {
    err := c.Send(&packet.ClientboundStatusPing{Time: p.(packet.ServerboundStatusPing).Time})
    if err != nil {
        log.Println(err)
    }
    c.Disconnect = true
}

type ServerListResponse struct {
    Version     GameVersion       `json:"version"`
    Players     ServerListPlayers `json:"players"`
    Description mc.Chat           `json:"description,omitempty"`
    Favicon     string            `json:"favicon,omitempty"`
}

type GameVersion struct {
    Name     string `json:"name"`
    Protocol int    `json:"protocol"`
}

type ServerListPlayers struct {
    Max    int                `json:"max"`
    Online int                `json:"online"`
    Sample []ServerListPlayer `json:"sample,omitempty"`
}

type ServerListPlayer struct {
    Name string    `json:"name"`
    Id   uuid.UUID `json:"id"`
}
