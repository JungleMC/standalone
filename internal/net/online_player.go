package net

import (
	"log"
	"sync"
	"time"

	"github.com/junglemc/JungleTree/internal/configuration"
	"github.com/junglemc/JungleTree/internal/net/protocol"
	"github.com/junglemc/JungleTree/internal/pkg/net/packets"
	"github.com/junglemc/JungleTree/pkg/chat"
	"github.com/junglemc/JungleTree/pkg/entity"
	"github.com/junglemc/JungleTree/pkg/inventory"
	"github.com/junglemc/JungleTree/pkg/util"
)

var onlinePlayers = make([]OnlinePlayer, 0)
var wait = &sync.WaitGroup{}

type OnlinePlayer struct {
	Client            *Client              `json:"-"`
	Entity            *entity.LivingEntity `json:"-"`
	ClientBrand       string               `json:"-"`
	Gamemode          util.GameMode
	Difficulty        util.Difficulty
	Locale            string
	ViewDistance      byte
	ChatMode          *chat.Mode
	ChatColorsEnabled bool
	SkinParts         byte
	MainHand          *util.Hand
	Inventory         inventory.Player `json:"-"`
	Hotbar            inventory.Hotbar `json:"-"`
}

func (o OnlinePlayer) String() string {
	return o.Client.Profile.Name
}

func tick(c *Client, time time.Time) (err error) {
	return
}

func Connect(c *Client) {
	if _, player, ok := getOnlinePlayer(c); ok {
		player.Client.Disconnect("You logged in from another location!")
	}

	playerEntity := entity.NewLivingEntity(entity.ByName("player"), c.Profile.ID, func(e *entity.LivingEntity, time time.Time) error {
		return tick(c, time)
	})

	player := OnlinePlayer{
		Client:     c,
		Entity:     playerEntity,
		Gamemode:   util.GameModeByName(configuration.Config().Gamemode),
		Difficulty: util.DifficultyByName(configuration.Config().Difficulty),
		Inventory:  inventory.Player{},
		Hotbar:     inventory.Hotbar{},
	}
	wait.Wait()
	onlinePlayers = append(onlinePlayers, player)
}

func Disconnect(c *Client, reason string) {
	if i, _, ok := getOnlinePlayer(c); ok {
		wait.Wait()
		if i+1 >= len(onlinePlayers) {
			onlinePlayers = onlinePlayers[:i]
		} else {
			onlinePlayers = append(onlinePlayers[:i], onlinePlayers[i+1:]...)
		}

		if reason != "" {
			log.Printf("%s disconnected: %s", c.Profile.Name, reason)
			if c.Protocol == protocol.Login {
				_ = c.Send(&packets.ClientboundLoginDisconnectPacket{Reason: &chat.Message{Text: reason}})
			} else if c.Protocol == protocol.Play {
				_ = c.Send(&packets.ClientboundPlayKickDisconnect{Reason: &chat.Message{Text: reason}})
			}
		}
	}
}

func GetOnlinePlayers() int {
	return len(onlinePlayers)
}

func GetOnlinePlayer(c *Client) (p *OnlinePlayer, ok bool) {
	_, p, ok = getOnlinePlayer(c)
	return
}

func getOnlinePlayer(c *Client) (index int, p *OnlinePlayer, ok bool) {
	wait.Add(1)
	defer wait.Done()
	for i, o := range onlinePlayers {
		if o.Client.Profile.ID == c.Profile.ID {
			return i, &onlinePlayers[i], true
		}
	}
	return 0, nil, false
}
