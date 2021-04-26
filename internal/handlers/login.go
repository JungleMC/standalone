package handlers

import (
    "bufio"
    "bytes"
    "github.com/google/uuid"
    "github.com/junglemc/JungleTree/internal/player"
    "github.com/junglemc/JungleTree/pkg"
    "github.com/junglemc/mc"
    "github.com/junglemc/mc/ability"
    "github.com/junglemc/net"
    "github.com/junglemc/net/auth"
    "github.com/junglemc/net/codec"
    "github.com/junglemc/net/codec/primitives"
    "github.com/junglemc/net/packet"
    "github.com/junglemc/world"
    "github.com/junglemc/world/dimensions"
)

func loginStart(c *net.Client, p codec.Packet) (err error) {
    pkt := p.(packet.ServerboundLoginStartPacket)

    c.Profile.Name = pkt.Username

    if c.Server.OnlineMode {
        pkt := &packet.ClientboundLoginEncryptionRequest{
            ServerId:    "",
            PublicKey:   c.Server.PublicKey(),
            VerifyToken: c.ExpectedVerifyToken,
        }
        return c.Send(pkt)
    } else {
        if c.Server.CompressionThreshold > 0 {
            err = c.EnableCompression()
            if err != nil {
                return
            }
        }

        c.Profile.ID, _ = uuid.NewRandom()
        err = c.LoginSuccess()
        if err != nil {
            return
        }
        return joinGame(c)
    }
}

func loginEncryptionResponse(c *net.Client, p codec.Packet) (err error) {
    pkt := p.(packet.ServerboundLoginEncryptionResponsePacket)

    sharedSecret, err := auth.DecryptLoginResponse(c.Server.PrivateKey(), c.Server.PublicKey(), c.ExpectedVerifyToken, pkt.VerifyToken, pkt.SharedSecret, &c.Profile)
    if err != nil {
        return
    }

    err = c.EnableEncryption(sharedSecret)
    if err != nil {
        return
    }

    if c.Server.CompressionThreshold > 0 {
        err = c.EnableCompression()
        if err != nil {
            return
        }
    }

    err = c.LoginSuccess()
    if err != nil {
        return
    }
    return joinGame(c)
}

func joinGame(c *net.Client) (err error) {
    err = sendJoinGame(c)
    if err != nil {
        return
    }

    player.Connect(c)

    err = sendServerBrand(c)
    if err != nil {
        return
    }

    err = sendServerDifficulty(c)
    if err != nil {
        return
    }

    return sendClientAbilities(c)
}

func sendJoinGame(c *net.Client) (err error) {
    join := &packet.ClientboundJoinGamePacket{
        EntityId:            0,
        IsHardcore:          false,
        GameMode:            mc.Survival,
        PreviousGameMode:    -1,
        WorldNames:          []string{"world"},
        DimensionCodec:      world.Codec(),
        Dimension:           *dimensions.ByName("overworld"),
        WorldName:           "world",
        HashedSeed:          0,
        MaxPlayers:          10,
        ViewDistance:        50,
        ReducedDebugInfo:    false,
        EnableRespawnScreen: true,
        IsDebug:             true,
        IsFlat:              false,
    }
    return c.Send(join)
}

func sendServerBrand(c *net.Client) (err error) {
    b := &bytes.Buffer{}
    buf := bufio.NewWriter(b)
    err = primitives.WriteString(buf, pkg.Brand)
    if err != nil {
        return
    }

    pkt := &packet.ClientboundPluginMessagePacket{
        Channel: mc.Identifier{
            Prefix: "minecraft",
            Name:   "brand",
        },
        Data: b.Bytes(),
    }
    return c.Send(pkt)
}

func sendServerDifficulty(c *net.Client) (err error) {
    pkt := &packet.ClientboundServerDifficultyPacket{
        Difficulty:       pkg.Config().Difficulty.Byte(),
        DifficultyLocked: false,
    }
    return c.Send(pkt)
}

func sendClientAbilities(c *net.Client) (err error) {
    onlinePlayer, ok := player.GetOnlinePlayer(c)
    if !ok {
        return
    }

    abilities := ability.Set(0, ability.Invulnerable)
    if onlinePlayer.Gamemode == mc.Creative {
        abilities = ability.Set(abilities, ability.CreativeMode)
    }

    pkt := &packet.ClientboundPlayerAbilitiesPacket{
        Flags:        byte(abilities),
        FlyingSpeed:  0.5,
        WalkingSpeed: 0.1,
    }
    return c.Send(pkt)
}
