package startup

import (
	"log"

	"github.com/junglemc/JungleTree/internal/storage"
	"github.com/junglemc/JungleTree/pkg/block"
	"github.com/junglemc/JungleTree/pkg/crafting"
	"github.com/junglemc/JungleTree/pkg/entity"
	"github.com/junglemc/JungleTree/pkg/event"
	"github.com/junglemc/JungleTree/pkg/item"
	"github.com/junglemc/JungleTree/pkg/world/biome"
	"github.com/junglemc/JungleTree/pkg/world/dimensions"
)

const (
	TicksPerSecond = 20
)

func Init() {
	event.Trigger(event.ServerStartupEvent{})
	loadStorage()
	loadDimensions()
	loadBiomes()
	loadBlocks()
	loadItems()
	loadRecipes()
	loadEntities()
}

func loadStorage() {
	log.Println("\t* Loading LevelDB")
	err := storage.Load()
	if err != nil {
		log.Panicln(err)
	}
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
