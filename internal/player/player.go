package player

import (
    "github.com/junglemc/JungleTree/pkg"
    "github.com/junglemc/entity"
    "github.com/junglemc/mc"
    "github.com/junglemc/mc/chat"
    "github.com/junglemc/net"
    "github.com/junglemc/net/codec"
    "github.com/junglemc/net/packet"
    "log"
    "sync"
    "time"
)

var onlinePlayers = make([]OnlinePlayer, 0)
var wait = &sync.WaitGroup{}

type OnlinePlayer struct {
    Client            *net.Client          `json:"-"`
    Entity            *entity.LivingEntity `json:"-"`
    ClientBrand       string               `json:"-"`
    Gamemode          mc.GameMode
    Difficulty        mc.Difficulty
    Locale            string
    ViewDistance      byte
    ChatMode          chat.Mode
    ChatColorsEnabled bool
    SkinParts         byte
    MainHand          mc.Hand
}

func (o OnlinePlayer) String() string {
    return o.Client.Profile.Name
}

func tick(c *net.Client, time time.Time) (err error) {
    return
}

func Connect(c *net.Client) {

    if _, player, ok := getOnlinePlayer(c); ok {
        player.Client.Disconnect(&chat.Message{Text: "You logged in from another location!"})
    }

    playerEntity := entity.NewLivingEntity(entity.ByName("player"), c.Profile.ID, func(e *entity.LivingEntity, time time.Time) error {
        return tick(c, time)
    })

    player := OnlinePlayer{
        Client:     c,
        Entity:     playerEntity,
        Gamemode:   pkg.Config().Gamemode,
        Difficulty: pkg.Config().Difficulty,
    }
    wait.Wait()
    onlinePlayers = append(onlinePlayers, player)
}

func Disconnect(c *net.Client, reason *chat.Message) {
    if i, _, ok := getOnlinePlayer(c); ok {
        wait.Wait()
        if i+1 >= len(onlinePlayers) {
            onlinePlayers = onlinePlayers[:i]
        } else {
            onlinePlayers = append(onlinePlayers[:i], onlinePlayers[i+1:]...)
        }

        if reason != nil {
            log.Printf("%s disconnected: %s", c.Profile.Name, reason.String())
            if c.Protocol == codec.ProtocolLogin {
                _ = c.Send(&packet.ClientboundLoginDisconnectPacket{Reason: *reason})
            } else if c.Protocol == codec.ProtocolPlay {
                _ = c.Send(&packet.ClientboundPlayKickDisconnect{Reason: *reason})
            }
        }
    }
}

func GetOnlinePlayers() int {
    return len(onlinePlayers)
}

func GetOnlinePlayer(c *net.Client) (p *OnlinePlayer, ok bool) {
    _, p, ok = getOnlinePlayer(c)
    return
}

func getOnlinePlayer(c *net.Client) (index int, p *OnlinePlayer, ok bool) {
    wait.Add(1)
    defer wait.Done()
    for i, o := range onlinePlayers {
        if o.Client.Profile.ID == c.Profile.ID {
            return i, &onlinePlayers[i], true
        }
    }
    return 0, nil, false
}
