package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/junglemc/JungleTree/internal/player"
	"github.com/junglemc/JungleTree/pkg"
	"github.com/junglemc/net"
	"github.com/junglemc/net/codec"
	packet2 "github.com/junglemc/packet"
	"log"
)

func playPluginMessage(c *net.Client, p net.Packet) (err error) {
	pkt := p.(packet2.ServerboundPluginMessagePacket)

	if pkt.Channel.Prefix() == "minecraft" && pkt.Channel.Name() == "brand" {
		buf := bytes.NewBuffer(pkt.Data)
		brand := codec.ReadString(buf)

		if onlinePlayer, ok := player.GetOnlinePlayer(c); ok {
			onlinePlayer.ClientBrand = brand

			if pkg.Config().DebugMode {
				log.Printf("Client brand for %s: %s\n", c.Profile.Name, onlinePlayer.ClientBrand)
			}
		}
	}
	return
}

func playClientSettings(c *net.Client, p net.Packet) (err error) {
	pkt := p.(packet2.ServerboundClientSettingsPacket)

	onlinePlayer, ok := player.GetOnlinePlayer(c)
	if !ok {
		return
	}
	onlinePlayer.Locale = pkt.Locale
	onlinePlayer.ViewDistance = pkt.ViewDistance
	onlinePlayer.ChatMode = pkt.ChatMode
	onlinePlayer.ChatColorsEnabled = pkt.ChatColorsEnabled
	onlinePlayer.MainHand = pkt.MainHand

	if pkg.Config().DebugMode {
		data, _ := json.MarshalIndent(onlinePlayer, "", "  ")
		log.Printf("Client settings for %s:\n%s\n", c.Profile.Name, string(data))
	}

	itemChange := &packet2.ClientboundHeldItemChangePacket{
		Slot: byte(onlinePlayer.Hotbar.SelectedIndex),
	}
	return c.Send(itemChange)
}
