package main

import (
    _ "embed"
    "github.com/junglemc/JungleTree/internal/handlers"
    "github.com/junglemc/entity"
    "github.com/junglemc/net"
    "github.com/junglemc/world/biomes"
    "github.com/junglemc/world/blocks"
    "github.com/junglemc/world/dimensions"
    "log"
)

const (
    JungleTreeVersion = "0.0.5" // TODO: Load from git or tags?
    thinLine          = "------------------------------------"
    thickLine         = "===================================="
    TicksPerSecond    = 20
)

func main() {
    s := net.NewServer("0.0.0.0", 25565, true, 0, false, handlers.Handshake, handlers.Status, handlers.Login, handlers.Play)

    log.Println(thickLine)
    log.Println("Starting JungleTree Server v" + JungleTreeVersion)
    log.Println(thickLine)

    loadDimensions()
    loadBiomes()
    loadBlocks()
    loadEntities()

    log.Println("Done!")
    log.Println(thinLine)

    log.Printf("Server listening on: %s:%d", s.Address, s.Port)
    err := s.Listen()
    if err != nil {
        log.Panicln(err)
    }
}

func loadBlocks() {
    log.Println("\t* Loading blocks")

    err := blocks.Load()
    if err != nil {
        log.Panicln(err)
    }
}

func loadBiomes() {
    log.Println("\t* Loading biomes")

    err := biomes.Load()
    if err != nil {
        log.Panicln(err)
    }
}

func loadDimensions() {
    log.Println("\t* Loading dimensions")
    err := dimensions.Load()
    if err != nil {
        log.Panicln(err)
    }
}

func loadEntities() {
    log.Println("\t* Loading entities")
    err := entity.Load()
    if err != nil {
        log.Panicln(err)
    }

    entityThread := entity.EntityRunner{TPS: TicksPerSecond}
    entityThread.Run()
}
