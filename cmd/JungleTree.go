package main

import (
    "github.com/junglemc/JungleTree/internal/handlers"
    "github.com/junglemc/net"
    "log"
)

func main() {
    s := net.NewServer("0.0.0.0", 25565, true, 256, true, handlers.Handshake, handlers.Status, handlers.Login, nil)

    err := s.Listen()
    if err != nil {
        log.Panicln(err)
    }
}
