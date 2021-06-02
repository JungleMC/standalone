package handlers

import (
	"encoding/json"
	"github.com/junglemc/JungleTree/internal/configuration"
	. "github.com/junglemc/JungleTree/internal/net"
	. "github.com/junglemc/JungleTree/internal/pkg/net/packets"
	"github.com/junglemc/JungleTree/pkg/chat"
)

func statusRequest(c *Client, _ Packet) (err error) {
	response := ServerListResponse{
		Version: GameVersion{
			Name:     "1.16.5",
			Protocol: 754,
		},
		Players: ServerListPlayers{
			Max:    configuration.Config().MaxOnlinePlayers,
			Online: GetOnlinePlayers(),
			Sample: []ServerListPlayer{},
		},
		Description: &chat.Message{Text: configuration.Config().MOTD},
	}

	data, err := json.Marshal(response)
	if err != nil {
		return
	}

	responsePkt := &ClientboundStatusResponsePacket{Response: string(data)}
	return c.Send(responsePkt)
}

func statusPing(c *Client, p Packet) (err error) {
	return c.Send(&ClientboundStatusPongPacket{Time: p.(ServerboundStatusPingPacket).Time})
}
