package main

import (
    "github.com/junglemc/JungleTree/internal/handlers"
    "github.com/junglemc/net"
    "log"
)

func main() {
    s := net.NewServer("0.0.0.0", 25565, handlers.Handshake, handlers.Status, handlers.Login, nil)

    s.GenerateKeys()

    err := s.Listen()
    if err != nil {
        log.Fatalln(err)
    }
}