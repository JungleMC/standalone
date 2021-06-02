package player

import (
	"github.com/junglemc/JungleTree"
	"github.com/junglemc/JungleTree/chat"
	"github.com/junglemc/JungleTree/entity"
	"github.com/junglemc/JungleTree/inventory"
	"github.com/junglemc/JungleTree/net"
	"github.com/junglemc/JungleTree/net/protocol"
	. "github.com/junglemc/JungleTree/packet"
	"github.com/junglemc/JungleTree/pkg"
	"log"
	"sync"
	"time"
)

var onlinePlayers = make([]OnlinePlayer, 0)
var wait = &sync.WaitGroup{}

type OnlinePlayer struct {
	Client            *net.Client          `json:"-"`
	Entity            *entity.LivingEntity `json:"-"`
	ClientBrand       string                `json:"-"`
	Gamemode          JungleTree.GameMode
	Difficulty        JungleTree.Difficulty
	Locale            string
	ViewDistance      byte
	ChatMode          *chat.Mode
	ChatColorsEnabled bool
	SkinParts         byte
	MainHand          *JungleTree.Hand
	Inventory         inventory.Player `json:"-"`
	Hotbar            inventory.Hotbar `json:"-"`
}

func (o OnlinePlayer) String() string {
	return o.Client.Profile.Name
}

func tick(c *net.Client, time time.Time) (err error) {
	return
}

func Connect(c *net.Client) {
	if _, player, ok := getOnlinePlayer(c); ok {
		player.Client.Disconnect("You logged in from another location!")
	}

	playerEntity := entity.NewLivingEntity(entity.ByName("player"), c.Profile.ID, func(e *entity.LivingEntity, time time.Time) error {
		return tick(c, time)
	})

	player := OnlinePlayer{
		Client:     c,
		Entity:     playerEntity,
		Gamemode:   JungleTree.GameModeByName(pkg.Config().Gamemode),
		Difficulty: JungleTree.DifficultyByName(pkg.Config().Difficulty),
		Inventory:  inventory.Player{},
		Hotbar:     inventory.Hotbar{},
	}
	wait.Wait()
	onlinePlayers = append(onlinePlayers, player)
}

func Disconnect(c *net.Client, reason string) {
	if i, _, ok := getOnlinePlayer(c); ok {
		wait.Wait()
		if i+1 >= len(onlinePlayers) {
			onlinePlayers = onlinePlayers[:i]
		} else {
			onlinePlayers = append(onlinePlayers[:i], onlinePlayers[i+1:]...)
		}

		if reason != "" {
			log.Printf("%s disconnected: %s", c.Profile.Name, reason)
			if c.Protocol == protocol.ProtocolLogin {
				_ = c.Send(&ClientboundLoginDisconnectPacket{Reason: chat.Message{Text: reason}})
			} else if c.Protocol == protocol.ProtocolPlay {
				_ = c.Send(&ClientboundPlayKickDisconnect{Reason: chat.Message{Text: reason}})
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
