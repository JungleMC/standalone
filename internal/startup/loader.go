package startup

import (
    "github.com/junglemc/JungleTree/pkg"
    "github.com/junglemc/JungleTree/pkg/block"
    "github.com/junglemc/JungleTree/pkg/crafting"
    "github.com/junglemc/JungleTree/pkg/entity"
    "github.com/junglemc/JungleTree/pkg/event"
    "github.com/junglemc/JungleTree/pkg/item"
    "github.com/junglemc/JungleTree/pkg/world/biome"
    "github.com/junglemc/JungleTree/pkg/world/dimensions"
    "log"
)

const (
    thinLine       = "------------------------------------"
    thickLine      = "===================================="
    TicksPerSecond = 20
)

func Load() {
    log.Println(thickLine)
    log.Println("Starting JungleTree Server v" + pkg.Version)
    log.Println(thickLine)

    loadDimensions()
    loadBiomes()
    loadBlocks()
    loadItems()
    loadRecipes()
    loadEntities()

    log.Println("Done!")
    log.Println(thinLine)

    event.Trigger(event.ServerLoadEvent{})
}

func loadBlocks() {
    log.Println("\t* Loading blocks")

    err := block.Load()
    if err != nil {
        log.Panicln(err)
    }
}

func loadBiomes() {
    log.Println("\t* Loading biomes")

    err := biome.Load()
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

    entityThread := entity.Runner{TPS: TicksPerSecond}
    entityThread.Run()
}

func loadItems() {
    log.Println("\t* Loading items")
    err := item.Load()
    if err != nil {
        log.Panicln(err)
    }
}

func loadRecipes() {
    log.Println("\t* Loading recipes")
    err := crafting.Load()
    if err != nil {
        log.Panicln(err)
    }
}
