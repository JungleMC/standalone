package handlers

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/junglemc/JungleTree/internal/configuration"
	. "github.com/junglemc/JungleTree/internal/net"
	. "github.com/junglemc/JungleTree/internal/pkg/net/packets"
	. "github.com/junglemc/JungleTree/pkg/codec"
)

func playPluginMessage(c *Client, p Packet) (err error) {
	pkt := p.(ServerboundPluginMessagePacket)

	if pkt.Channel.Prefix() == "minecraft" && pkt.Channel.Name() == "brand" {
		buf := bytes.NewBuffer(pkt.Data)
		brand := ReadString(buf)

		if onlinePlayer, ok := GetOnlinePlayer(c); ok {
			onlinePlayer.ClientBrand = brand

			if configuration.Config().DebugMode {
				log.Printf("Client brand for %s: %s\n", c.Profile.Name, onlinePlayer.ClientBrand)
			}
		}
	}
	return
}

func playClientSettings(c *Client, p Packet) (err error) {
	pkt := p.(ServerboundClientSettingsPacket)

	onlinePlayer, ok := GetOnlinePlayer(c)
	if !ok {
		return
	}
	onlinePlayer.Locale = pkt.Locale
	onlinePlayer.ViewDistance = pkt.ViewDistance
	onlinePlayer.ChatMode = pkt.ChatMode
	onlinePlayer.ChatColorsEnabled = pkt.ChatColorsEnabled
	onlinePlayer.MainHand = pkt.MainHand

	if configuration.Config().DebugMode {
		data, _ := json.MarshalIndent(onlinePlayer, "", "  ")
		log.Printf("Client settings for %s:\n%s\n", c.Profile.Name, string(data))
	}

	itemChange := &ClientboundHeldItemChangePacket{
		Slot: byte(onlinePlayer.Hotbar.SelectedIndex),
	}
	return c.Send(itemChange)
}
