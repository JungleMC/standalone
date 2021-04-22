package internal

import (
    "github.com/junglemc/entity"
    "github.com/junglemc/net"
)

type OnlinePlayer struct {
    Client *net.Client
    Entity *entity.LivingEntity
}
