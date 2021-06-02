package handlers

import (
	"encoding/json"
	"github.com/junglemc/JungleTree/chat"
	"github.com/junglemc/JungleTree/internal/player"
	"github.com/junglemc/JungleTree/net"
	. "github.com/junglemc/JungleTree/packet"
	"github.com/junglemc/JungleTree/pkg"
	. "github.com/junglemc/JungleTree/status"
)

func statusRequest(c *net.Client, p net.Packet) (err error) {
	response := ServerListResponse{
		Version: GameVersion{
			Name:     "1.16.5",
			Protocol: 754,
		},
		Players: ServerListPlayers{
			Max:    pkg.Config().MaxOnlinePlayers,
			Online: player.GetOnlinePlayers(),
			Sample: []ServerListPlayer{},
		},
		Description: chat.Message{Text: pkg.Config().MOTD},
	}

	data, err := json.Marshal(response)
	if err != nil {
		return
	}

	responsePkt := &ClientboundStatusResponsePacket{Response: string(data)}
	return c.Send(responsePkt)
}

func statusPing(c *net.Client, p net.Packet) (err error) {
	return c.Send(&ClientboundStatusPongPacket{Time: p.(ServerboundStatusPingPacket).Time})
}
