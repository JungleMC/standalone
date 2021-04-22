package main

import (
    _ "embed"
    "github.com/junglemc/JungleTree/internal/handlers"
    "github.com/junglemc/net"
    "github.com/junglemc/world/biomes"
    "github.com/junglemc/world/blocks"
    "github.com/junglemc/world/dimensions"
    "log"
)

const (
    JungleTreeVersion = "0.0.3" // TODO: Load from git or tags?
    thinLine          = "------------------------------------"
    thickLine         = "===================================="
)

func main() {
    s := net.NewServer("0.0.0.0", 25565, true, 0, true, handlers.Handshake, handlers.Status, handlers.Login, nil)

    log.Println(thickLine)
    log.Println("Starting JungleTree Server v" + JungleTreeVersion)
    log.Println(thickLine)

    loadBlocks(s)
    loadBiomes(s)
    loadDimensions(s)

    if s.Debug {
        log.Println("Done!")
        log.Println(thinLine)
    }

    log.Printf("Server listening on: %s:%d", s.Address, s.Port)
    err := s.Listen()
    if err != nil {
        log.Panicln(err)
    }
}

func loadBlocks(s *net.Server) {
    if s.Debug {
        log.Println("\t* Loading blocks")
    }

    err := blocks.Load()
    if err != nil {
        log.Panicln(err)
    }
}

func loadBiomes(s *net.Server) {
    if s.Debug {
        log.Println("\t* Loading biomes")
    }

    err := biomes.Load()
    if err != nil {
        log.Panicln(err)
    }
}

func loadDimensions(s *net.Server) {
    if s.Debug {
        log.Println("\t* Loading dimensions")
    }

    err := dimensions.Load()
    if err != nil {
        log.Panicln(err)
    }
}
