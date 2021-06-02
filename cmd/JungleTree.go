package main

import (
	"github.com/junglemc/JungleTree/crafting"
	"github.com/junglemc/JungleTree/entity"
	"github.com/junglemc/JungleTree/internal/handlers"
	"github.com/junglemc/JungleTree/internal/player"
	"github.com/junglemc/JungleTree/item"
	"github.com/junglemc/JungleTree/net"
	"github.com/junglemc/JungleTree/pkg"
	"github.com/junglemc/JungleTree/world/biomes"
	"github.com/junglemc/JungleTree/world/blocks"
	"github.com/junglemc/JungleTree/world/dimensions"
	"log"
)

const (
	JungleTreeVersion = "0.0.8" // TODO: Load from git or tags?
	thinLine          = "------------------------------------"
	thickLine         = "===================================="
	TicksPerSecond    = 20
)

func main() {
	conf := pkg.Config()

	s := net.NewServer(
		conf.Network.IP, conf.Network.Port, conf.JavaEdition.OnlineMode, conf.Network.NetworkCompressionThreshold,
		conf.DebugMode, conf.Verbose, handlers.Handshake, handlers.Status, handlers.Login, handlers.Play, player.Disconnect,
	)

	log.Println(thickLine)
	log.Println("Starting JungleTree Server v" + JungleTreeVersion)
	log.Println(thickLine)

	loadDimensions()
	loadBiomes()
	loadBlocks()
	loadItems()
	loadRecipes()
	loadEntities()

	log.Println("Done!")
	log.Println(thinLine)

	addr := s.Address
	if addr == "" {
		addr = "*"
	}

	log.Printf("Server listening on: %s:%d", addr, s.Port)
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
