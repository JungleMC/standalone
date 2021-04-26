package handlers

import (
    "bufio"
    "bytes"
    "encoding/json"
    "github.com/junglemc/JungleTree/internal/player"
    "github.com/junglemc/JungleTree/pkg"
    "github.com/junglemc/net"
    "github.com/junglemc/net/codec"
    "github.com/junglemc/net/codec/primitives"
    "github.com/junglemc/net/packet"
    "log"
)

func playPluginMessage(c *net.Client, p codec.Packet) (err error) {
    pkt := p.(packet.ServerboundPluginMessagePacket)

    if pkt.Channel.Prefix == "minecraft" && pkt.Channel.Name == "brand" {
        buf := bufio.NewReader(bytes.NewReader(pkt.Data))
        brand := ""
        brand, err = primitives.ReadString(buf)
        if err != nil {
            return
        }

        if onlinePlayer, ok := player.GetOnlinePlayer(c); ok {
            onlinePlayer.ClientBrand = brand

            if pkg.Config().DebugMode {
                log.Printf("Client brand for %s: %s\n", c.Profile.Name, onlinePlayer.ClientBrand)
            }
        }
    }
    return
}

func playClientSettings(c *net.Client, p codec.Packet) (err error) {
    pkt := p.(packet.ServerboundClientSettingsPacket)

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

    itemChange := &packet.ClientboundHeldItemChangePacket{
        Slot: byte(onlinePlayer.Hotbar.SelectedIndex),
    }
    return c.Send(itemChange)
}
