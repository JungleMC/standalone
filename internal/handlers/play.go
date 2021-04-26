package handlers

import (
    "bufio"
    "bytes"
    "github.com/junglemc/JungleTree/internal/player"
    "github.com/junglemc/net"
    "github.com/junglemc/net/codec"
    "github.com/junglemc/net/codec/primitives"
    "github.com/junglemc/net/packet"
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
        }
    }
    return
}

func playClientSettings(c *net.Client, p codec.Packet) (err error) {
    // pkt := p.(packet.ServerboundClientSettingsPacket)
    return
}
