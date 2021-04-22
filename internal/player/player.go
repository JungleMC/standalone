package player

import (
    "errors"
    "github.com/junglemc/entity"
    "github.com/junglemc/net"
    "sync"
    "time"
)

var onlinePlayers = make([]OnlinePlayer, 0)
var wait = &sync.WaitGroup{}

type OnlinePlayer struct {
    Client      *net.Client
    Entity      *entity.LivingEntity
    ClientBrand string
}

func (o OnlinePlayer) String() string {
    return o.Client.Profile.Name
}

func tick(c *net.Client, time time.Time) (err error) {
    return
}

func Connect(c *net.Client) {
    wait.Add(1)

    playerEntity := entity.NewLivingEntity(entity.ByName("player"), c.Profile.ID, func(e *entity.LivingEntity, time time.Time) error {
        return tick(c, time)
    })

    player := OnlinePlayer{
        Client: c,
        Entity: playerEntity,
    }
    onlinePlayers = append(onlinePlayers, player)
    wait.Done()
}

func GetOnlinePlayer(c *net.Client) *OnlinePlayer {
    wait.Add(1)
    for i, o := range onlinePlayers {
        if o.Client.Profile.ID == c.Profile.ID {
            return &onlinePlayers[i]
        }
    }
    wait.Done()
    panic(errors.New("online player not found for client: " + c.Profile.Name))
}
