package main

import (
    "github.com/junglemc/JungleTree/internal/handlers"
    "github.com/junglemc/net"
    "github.com/junglemc/world"
    "log"
    "path/filepath"
)

const (
    JungleTreeVersion = "0.0.3" // TODO: Load from git or tags?
    MinecraftVersion  = "1.16.5"
    thinLine          = "------------------------------------"
    thickLine         = "===================================="
)

func main() {
    s := net.NewServer("0.0.0.0", 25565, true, 256, true, handlers.Handshake, handlers.Status, handlers.Login, nil)

    log.Println(thickLine)
    log.Println("Starting JungleTree Server v" + JungleTreeVersion)
    log.Println(thickLine)

    if s.Debug {
        log.Println("\t* Loading blocks")
    }

    err := world.LoadBlocks(filepath.Join("configs", MinecraftVersion, "blocks.json"))
    if err != nil {
        log.Panicln(err)
    }

    if s.Debug {
        log.Println("Done!")
        log.Println(thinLine)
    }

    log.Printf("Server listening on: %s:%d", s.Address, s.Port)
    err = s.Listen()
    if err != nil {
        log.Panicln(err)
    }
}
